# VK\_KHR\_ray\_tracing\_maintenance1(3) Manual Page

## Name

VK\_KHR\_ray\_tracing\_maintenance1 - device extension



## [](#_registered_extension_number)Registered Extension Number

387

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_EXT\_device\_generated\_commands
- Interacts with VK\_KHR\_ray\_tracing\_pipeline
- Interacts with VK\_KHR\_synchronization2

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_ray\_cull\_mask](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_ray_cull_mask.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_ray_tracing_maintenance1%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_ray_tracing_maintenance1%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-02-21

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_ray_cull_mask`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_ray_cull_mask.txt)
- Interacts with `VK_KHR_ray_tracing_pipeline`
- Interacts with `VK_KHR_synchronization2`

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

## [](#_description)Description

`VK_KHR_ray_tracing_maintenance1` adds a collection of minor ray tracing features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Adds support for the `SPV_KHR_ray_cull_mask` SPIR-V extension in Vulkan. This extension provides access to built-in `CullMaskKHR` shader variable which contains the value of the `OpTrace*` `Cull Mask` parameter. This new shader variable is accessible in the intersection, any-hit, closest hit and miss shader stages.
- Adds support for a new pipeline stage and access mask built on top of `VK_KHR_synchronization2`:
  
  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` to specify execution of [acceleration structure copy commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-copying)
  - `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR` to specify read access to a [shader binding table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shader-binding-table) in any shader pipeline stage
- Adds two new acceleration structure query parameters:
  
  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` to query the acceleration structure size on the device timeline
  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR` to query the number of bottom level acceleration structure pointers for serialization
- Adds an optional new indirect ray tracing dispatch command, [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirect2KHR.html), which sources the shader binding table parameters as well as the dispatch dimensions from the device. The [`rayTracingPipelineTraceRaysIndirect2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTracingPipelineTraceRaysIndirect2) feature indicates whether this functionality is supported.

## [](#_new_commands)New Commands

If [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirect2KHR.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR.html)

If [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- [VkTraceRaysIndirectCommand2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTraceRaysIndirectCommand2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_RAY_TRACING_MAINTENANCE_1_EXTENSION_NAME`
- `VK_KHR_RAY_TRACING_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
  - `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MAINTENANCE_1_FEATURES_KHR`

If [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) or [Vulkan Version 1.3](#versions-1.3) and [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`

If [VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html) is supported:

- Extending [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html):
  
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`

If [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`

## [](#_new_built_in_variables)New Built-In Variables

- [`CullMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-cullmask)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`RayCullMaskKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayCullMaskKHR)

## [](#_issues)Issues

None Yet!

## [](#_version_history)Version History

- Revision 1, 2022-02-21 (Members of the Vulkan Ray Tracing TSG)
  
  - internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_ray_tracing_maintenance1).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0