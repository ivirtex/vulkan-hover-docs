# VK\_NV\_ray\_tracing\_motion\_blur(3) Manual Page

## Name

VK\_NV\_ray\_tracing\_motion\_blur - device extension



## [](#_registered_extension_number)Registered Extension Number

328

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_ray\_tracing\_motion\_blur](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_ray_tracing_motion_blur.html)

## [](#_contact)Contact

- Eric Werness

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-16

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_ray_tracing_motion_blur`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_ray_tracing_motion_blur.txt)

**Contributors**

- Eric Werness, NVIDIA
- Ashwin Lele, NVIDIA

## [](#_description)Description

Ray tracing support in the API provides an efficient mechanism to intersect rays against static geometry, but rendering algorithms often want to support motion, which is more efficiently supported with motion-specific algorithms. This extension adds a set of mechanisms to support fast tracing of moving geometry:

- A ray pipeline trace call which takes a time parameter
- Flags to enable motion support in an acceleration structure
- Support for time-varying vertex positions in a geometry
- Motion instances to move existing instances over time

The motion represented here is parameterized across a normalized timestep between 0.0 and 1.0. A motion trace using `OpTraceRayMotionNV` provides a time within that normalized range to be used when intersecting that ray with geometry. The geometry can be provided with motion by a combination of adding a second vertex position for time of 1.0 using `VkAccelerationStructureGeometryMotionTrianglesDataNV` and providing multiple transforms in the instance using `VkAccelerationStructureMotionInstanceNV`.

## [](#_new_structures)New Structures

- [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)
- [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html)
- [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)
- [VkSRTDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSRTDataNV.html)
- Extending [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html):
  
  - [VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInfoNV.html)
- Extending [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html):
  
  - [VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingMotionBlurFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingMotionBlurFeaturesNV.html)

## [](#_new_unions)New Unions

- [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceDataNV.html)

## [](#_new_enums)New Enums

- [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAccelerationStructureMotionInfoFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInfoFlagsNV.html)
- [VkAccelerationStructureMotionInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_RAY_TRACING_MOTION_BLUR_EXTENSION_NAME`
- `VK_NV_RAY_TRACING_MOTION_BLUR_SPEC_VERSION`
- Extending [VkAccelerationStructureCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateFlagBitsKHR.html):
  
  - `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV`
- Extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):
  
  - `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_MOTION_TRIANGLES_DATA_NV`
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MOTION_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MOTION_BLUR_FEATURES_NV`

## [](#_issues)Issues

(1) What size is VkAccelerationStructureMotionInstanceNV?

- Added a note on the structure size and made the stride explicit in the language.

(2) Allow arrayOfPointers for motion TLAS?

- Yes, with a packed encoding to minimize the amount of data sent for metadata.

## [](#_version_history)Version History

- Revision 1, 2020-06-16 (Eric Werness, Ashwin Lele)
  
  - Initial external release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_ray_tracing_motion_blur)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0