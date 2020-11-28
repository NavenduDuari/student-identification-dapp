pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Identity {

    struct AcademicDetailsI {
        string degreeId;
    }

    mapping(string => AcademicDetailsI) public studentIdVsAcademicDetails;

    function setAcademicDetails(string memory studentId, AcademicDetailsI memory academicDetails) public {
      studentIdVsAcademicDetails[studentId] = academicDetails;
    }

    function getAcademicDetails(string memory studentId) public view returns(AcademicDetailsI memory){
      return studentIdVsAcademicDetails[studentId];
   }
}