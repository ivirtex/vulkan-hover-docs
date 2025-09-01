# VK\_AMDX\_dense\_geometry\_format(3) Manual Page

## Name

VK\_AMDX\_dense\_geometry\_format - device extension



## [](#_registered_extension_number)Registered Extension Number

479

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)  
and  
     [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
     or  
     [Vulkan Version 1.4](#versions-1.4)

- **This is a *provisional* extension and must be used with caution. See the [description](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-provisional-header) of provisional header files for enablement and stability details.**

## [](#_contact)Contact

- Stu Smith [\[GitHub\]stu-s](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMDX_dense_geometry_format%5D%20%40stu-s%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMDX_dense_geometry_format%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_AMDX\_dense\_geometry\_format](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMDX_dense_geometry_format.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-07-10

**IP Status**

No known IP claims.

**Contributors**

- Stu Smith, AMD
- Josh Barczak, AMD
- Carsten Benthin, AMD
- David McAllister, AMD

## [](#_description)Description

This extension adds the ability to build ray tracing acceleration structures from pre-compressed `Dense Geometry Format` geometry data.

## [](#_new_structures)New Structures

- Extending [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html):
  
  - [VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX.html)

## [](#_new_enums)New Enums

- [VkCompressedTriangleFormatAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompressedTriangleFormatAMDX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMDX_DENSE_GEOMETRY_FORMAT_EXTENSION_NAME`
- `VK_AMDX_DENSE_GEOMETRY_FORMAT_SPEC_VERSION`
- `VK_COMPRESSED_TRIANGLE_FORMAT_DGF1_BYTE_ALIGNMENT_AMDX`
- `VK_COMPRESSED_TRIANGLE_FORMAT_DGF1_BYTE_STRIDE_AMDX`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_COMPRESSED_DATA_DGF1_BIT_AMDX`
- Extending [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html):
  
  - `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DENSE_GEOMETRY_FORMAT_TRIANGLES_DATA_AMDX`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DENSE_GEOMETRY_FORMAT_FEATURES_AMDX`

## [](#_issues)Issues

None.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2025-07-10 (Stu Smith)
  
  - Initial revision.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMDX_dense_geometry_format).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0