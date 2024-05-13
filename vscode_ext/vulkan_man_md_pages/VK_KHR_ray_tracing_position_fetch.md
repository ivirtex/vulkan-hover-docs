# VK_KHR_ray_tracing_position_fetch(3) Manual Page

## Name

VK_KHR_ray_tracing_position_fetch - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

482

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_ray_tracing_position_fetch](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_ray_tracing_position_fetch.html)

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_ray_tracing_position_fetch](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_ray_tracing_position_fetch.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-02-17

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_ray_tracing_position_fetch`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_tracing_position_fetch.txt)

- Interacts with
  [`VK_KHR_ray_tracing_pipeline`](VK_KHR_ray_tracing_pipeline.html)

- Interacts with [`VK_KHR_ray_query`](VK_KHR_ray_query.html)

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

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_ray_tracing_position_fetch` adds the ability to fetch the vertex
positions in the shader from a hit triangle as stored in the
acceleration structure.

An application adds
`VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR` to the
acceleration structure at build time. Then, if the hit is a triangle
geometry, the shader (any-hit or closest hit for ray pipelines or using
ray query) **can** fetch the three, three-component vertex positions in
object space, of the triangle which was hit.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_RAY_TRACING_POSITION_FETCH_EXTENSION_NAME`

- `VK_KHR_RAY_TRACING_POSITION_FETCH_SPEC_VERSION`

- Extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_POSITION_FETCH_FEATURES_KHR`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-hittrianglevertexpositions"
  target="_blank"
  rel="noopener"><code>HitTriangleVertexPositionsKHR</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayTracingPositionFetchKHR"
  target="_blank" rel="noopener">RayTracingPositionFetchKHR</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayQueryPositionFetchKHR"
  target="_blank" rel="noopener">RayQueryPositionFetchKHR</a>

## <a href="#_issues" class="anchor"></a>Issues

None Yet!

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-02-17 (Eric Werness)

  - internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_ray_tracing_position_fetch"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
