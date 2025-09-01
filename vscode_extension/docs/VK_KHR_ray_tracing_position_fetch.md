# VK\_KHR\_ray\_tracing\_position\_fetch(3) Manual Page

## Name

VK\_KHR\_ray\_tracing\_position\_fetch - device extension



## [](#_registered_extension_number)Registered Extension Number

482

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_ray\_tracing\_position\_fetch](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_ray_tracing_position_fetch.html)

## [](#_contact)Contact

- Eric Werness

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_ray\_tracing\_position\_fetch](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_ray_tracing_position_fetch.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-02-17

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_ray_tracing_position_fetch`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_tracing_position_fetch.txt)
- Interacts with `VK_KHR_ray_tracing_pipeline`
- Interacts with `VK_KHR_ray_query`

**Contributors**

- Eric Werness, NVIDIA
- Stu Smith, AMD
- Yuriy O’Donnell, Epic Games
- Ralph Potter, Samsung
- Joshua Barczak, Intel
- Lionel Landwerlin, Intel
- Andrew Garrard, Imagination Technologies
- Alex Bourd, Qualcomm
- Yunpeng Zhu, Huawei Technologies
- Marius Bjorge, Arm
- Daniel Koch, NVIDIA

## [](#_description)Description

`VK_KHR_ray_tracing_position_fetch` adds the ability to fetch the vertex positions in the shader from a hit triangle as stored in the acceleration structure.

An application adds `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_BIT_KHR` to the acceleration structure at build time. Then, if the hit is a triangle geometry, the shader (any-hit or closest hit for ray pipelines or using ray query) **can** fetch the three, three-component vertex positions in object space, of the triangle which was hit.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_RAY_TRACING_POSITION_FETCH_EXTENSION_NAME`
- `VK_KHR_RAY_TRACING_POSITION_FETCH_SPEC_VERSION`
- Extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):
  
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_BIT_KHR`
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_POSITION_FETCH_FEATURES_KHR`

## [](#_new_built_in_variables)New Built-In Variables

- [`HitTriangleVertexPositionsKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hittrianglevertexpositions)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [RayTracingPositionFetchKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingPositionFetchKHR)
- [RayQueryPositionFetchKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayQueryPositionFetchKHR)

## [](#_issues)Issues

None Yet!

## [](#_version_history)Version History

- Revision 1, 2023-02-17 (Eric Werness)
  
  - internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_ray_tracing_position_fetch).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0