# VK_NV_ray_tracing_motion_blur(3) Manual Page

## Name

VK_NV_ray_tracing_motion_blur - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

328

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_ray_tracing_motion_blur](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_ray_tracing_motion_blur.html)

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-16

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_ray_tracing_motion_blur`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_ray_tracing_motion_blur.txt)

**Contributors**  
- Eric Werness, NVIDIA

- Ashwin Lele, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Ray tracing support in the API provides an efficient mechanism to
intersect rays against static geometry, but rendering algorithms often
want to support motion, which is more efficiently supported with
motion-specific algorithms. This extension adds a set of mechanisms to
support fast tracing of moving geometry:

- A ray pipeline trace call which takes a time parameter

- Flags to enable motion support in an acceleration structure

- Support for time-varying vertex positions in a geometry

- Motion instances to move existing instances over time

The motion represented here is parameterized across a normalized
timestep between 0.0 and 1.0. A motion trace using `OpTraceRayMotionNV`
provides a time within that normalized range to be used when
intersecting that ray with geometry. The geometry can be provided with
motion by a combination of adding a second vertex position for time of
1.0 using `VkAccelerationStructureGeometryMotionTrianglesDataNV` and
providing multiple transforms in the instance using
`VkAccelerationStructureMotionInstanceNV`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)

- [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)

- [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)

- [VkSRTDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSRTDataNV.html)

- Extending
  [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html):

  - [VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoNV.html)

- Extending
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html):

  - [VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingMotionBlurFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingMotionBlurFeaturesNV.html)

## <a href="#_new_unions" class="anchor"></a>New Unions

- [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceDataNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkAccelerationStructureMotionInfoFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoFlagsNV.html)

- [VkAccelerationStructureMotionInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_RAY_TRACING_MOTION_BLUR_EXTENSION_NAME`

- `VK_NV_RAY_TRACING_MOTION_BLUR_SPEC_VERSION`

- Extending
  [VkAccelerationStructureCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagBitsKHR.html):

  - `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV`

- Extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):

  - `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_MOTION_TRIANGLES_DATA_NV`

  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MOTION_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MOTION_BLUR_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

\(1\) What size is VkAccelerationStructureMotionInstanceNV?

- Added a note on the structure size and made the stride explicit in the
  language.

\(2\) Allow arrayOfPointers for motion TLAS?

- Yes, with a packed encoding to minimize the amount of data sent for
  metadata.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-06-16 (Eric Werness, Ashwin Lele)

  - Initial external release

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_ray_tracing_motion_blur"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
