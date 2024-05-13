# VK_KHR_ray_tracing_maintenance1(3) Manual Page

## Name

VK_KHR_ray_tracing_maintenance1 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

387

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_ray_tracing_pipeline

- Interacts with VK_KHR_synchronization2

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_ray_cull_mask](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_ray_cull_mask.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_tracing_maintenance1%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_tracing_maintenance1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-02-21

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_ray_cull_mask`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_cull_mask.txt)

- Interacts with
  [`VK_KHR_ray_tracing_pipeline`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)

- Interacts with
  [`VK_KHR_synchronization2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)

**Contributors**  
- Stu Smith, AMD

- Tobias Hector, AMD

- Marius Bjorge, Arm

- Tom Olson, Arm

- Yuriy O’Donnell, Epic Games

- Yunpeng Zhu, Huawei

- Andrew Garrard, Imagination

- Dae Kim, Imagination

- Joshua Barczak, Intel

- Lionel Landwerlin, Intel

- Daniel Koch, NVIDIA

- Eric Werness, NVIDIA

- Spencer Fricke, Samsung

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_ray_tracing_maintenance1` adds a collection of minor ray tracing
features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Adds support for the `SPV_KHR_ray_cull_mask` SPIR-V extension in
  Vulkan. This extension provides access to built-in `CullMaskKHR`
  shader variable which contains the value of the `OpTrace*` `Cull Mask`
  parameter. This new shader variable is accessible in the intersection,
  any-hit, closest-hit and miss shader stages.

- Adds support for a new pipeline stage and access mask built on top of
  [`VK_KHR_synchronization2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html):

  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` to specify
    execution of <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-copying"
    target="_blank" rel="noopener">acceleration structure copy commands</a>

  - `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR` to specify read
    access to a <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-binding-table"
    target="_blank" rel="noopener">shader binding table</a> in any
    shader pipeline stage

- Adds two new acceleration structure query parameters:

  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` to query the
    acceleration structure size on the device timeline

  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
    to query the number of bottom level acceleration structure pointers
    for serialization

- Adds an optional new indirect ray tracing dispatch command,
  [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirect2KHR.html), which
  sources the shader binding table parameters as well as the dispatch
  dimensions from the device. The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rayTracingPipelineTraceRaysIndirect2"
  target="_blank"
  rel="noopener"><code>rayTracingPipelineTraceRaysIndirect2</code></a>
  feature indicates whether this functionality is supported.

## <a href="#_new_commands" class="anchor"></a>New Commands

If [VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html) is
supported:

- [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirect2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR.html)

If [VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html) is
supported:

- [VkTraceRaysIndirectCommand2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTraceRaysIndirectCommand2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_RAY_TRACING_MAINTENANCE_1_EXTENSION_NAME`

- `VK_KHR_RAY_TRACING_MAINTENANCE_1_SPEC_VERSION`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`

  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MAINTENANCE_1_FEATURES_KHR`

If [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html) or [Version
1.3](#versions-1.3) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-cullmask"
  target="_blank" rel="noopener"><code>CullMaskKHR</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-RayCullMaskKHR"
  target="_blank" rel="noopener"><code>RayCullMaskKHR</code></a>

## <a href="#_issues" class="anchor"></a>Issues

None Yet!

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-02-21 (Members of the Vulkan Ray Tracing TSG)

  - internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_ray_tracing_maintenance1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
