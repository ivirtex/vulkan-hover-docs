# VK\_EXT\_device\_generated\_commands(3) Manual Page

## Name

VK\_EXT\_device\_generated\_commands - device extension



## [](#_registered_extension_number)Registered Extension Number

573

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html)  
         or  
         [Vulkan Version 1.2](#versions-1.2)  
     and  
     [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_shader\_object

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_device_generated_commands%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_device_generated_commands%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_device\_generated\_commands](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_device_generated_commands.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-02-23

**Interactions and External Dependencies**

- This extension requires Vulkan 1.1
- This extension requires `VK_EXT_buffer_device_address` or `VK_KHR_buffer_device_address` or Vulkan 1.2 for the ability to bind vertex and index buffers on the device.
- This extension requires `VK_KHR_maintenance5` for the ability to use VkPipelineCreateFlags2KHR.
- This extension interacts with `VK_NV_mesh_shader`. If the latter extension is not supported, remove the command tokens to initiate NV mesh tasks drawing in this extension.
- This extension interacts with `VK_EXT_mesh_shader`. If the latter extension is not supported, remove the command tokens to initiate EXT mesh tasks drawing in this extension.
- This extension interacts with `VK_KHR_ray_tracing_pipeline`. If the latter extension is not supported, remove the command tokens to initiate ray tracing in this extension.
- This extension interacts with `VK_EXT_shader_object`. If the latter extension is not supported, remove references to shader objects in this extension.

**Contributors**

- Mike Blumenkrantz, VALVE
- Hans-Kristian Arntzen, VALVE
- Jan-Harald Fredriksen, ARM
- Spencer Fricke, LunarG
- Ricardo Garcia, Igalia
- Tobias Hector, AMD
- Baldur Karlsson, VALVE
- Christoph Kubisch, NVIDIA
- Lionel Landwerlin, INTEL
- Jon Leech, Khronos
- Ting Wei, ARM
- Ken Shanyi Zhang, AMD
- Faith Ekstrand, Collabora
- Vikram Kushwaha, NVIDIA
- Connor Abbott, VALVE
- Samuel Pitoiset, VALVE

## [](#_description)Description

This extension allows the device to generate a number of commands for command buffers. It provides a subset of functionality from both `VK_NV_device_generated_commands` and `VK_NV_device_generated_commands_compute` as well as some new features.

When rendering a large number of objects, the device can be leveraged to implement a number of critical functions, like updating matrices, or implementing occlusion culling, frustum culling, front to back sorting, etc. Implementing those on the device does not require any special extension, since an application is free to define its own data structures, and just process them using shaders.

To render objects which have been processed on the device, Vulkan has several ways to perform indirect rendering, from the most basic `vkCmdDrawIndirect` with one indirect draw to `vkCmdDrawIndirectCount` which supports multiple indirect draws batched together, with a way to determine number of draws at device execution time.

However, if rendering state needs to change between the indirect draws, then unextended Vulkan forces the application to speculatively record a prohibitive number of redundant indirect commands covering all possible state combinations - which could end up processing nothing after culling - or read back the processed stream and issue graphics command from the host. For very large scenes, the synchronization overhead and cost to generate the command buffer can become the bottleneck. This extension allows an application to generate a device side stream of state changes and commands, and convert it efficiently into a command buffer without having to read it back to the host.

Furthermore, it allows incremental changes to such command buffers by manipulating only partial sections of a command stream — for example pipeline and shader object bindings. Unextended Vulkan requires re-creation of entire command buffers in such a scenario, or updates synchronized on the host.

The intended usage for this extension is for the application to:

- create `VkBuffer` objects and retrieve physical addresses from them via [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- create a `VkIndirectExecutionSetEXT` for the ability to change shaders on the device.
- create a [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html), which lists the [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html) it wants to dynamically execute as an atomic command sequence. This step likely involves some internal device code compilation, since the intent is for the GPU to generate the command buffer based on the layout.
- fill the input stream buffers with the data for each of the inputs it needs. Each input is an array that will be filled with token-dependent data.
- set up a preprocess `VkBuffer` that uses memory according to the information retrieved via [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html).
- optionally preprocess the generated content using [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html), for example on an asynchronous compute queue, or for the purpose of reusing the data in multiple executions.
- call [vkCmdExecuteGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteGeneratedCommandsEXT.html) to create and execute the actual device commands for all sequences based on the inputs provided.

For each draw in a sequence, the following can be specified:

- a number of vertex buffer bindings
- a different index buffer, with an optional dynamic offset and index type
- a number of different push constants
- updates to bound shader stages

For each dispatch in a sequence, the following can be specified:

- a number of different push constants
- updates to bound shader stages

For each trace rays in a sequence, the following can be specified:

- a number of different push constants
- updates to bound shader stages

While the GPU can be faster than a CPU to generate the commands, it will not happen asynchronously to the device, therefore the primary use case is generating “less” total work (occlusion culling, classification to use specialized shaders, etc.).

## [](#_new_object_types)New Object Types

- [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html)
- [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html)

## [](#_new_commands)New Commands

- [vkCmdExecuteGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteGeneratedCommandsEXT.html)
- [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html)
- [vkCreateIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectCommandsLayoutEXT.html)
- [vkCreateIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectExecutionSetEXT.html)
- [vkDestroyIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyIndirectCommandsLayoutEXT.html)
- [vkDestroyIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyIndirectExecutionSetEXT.html)
- [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html)
- [vkUpdateIndirectExecutionSetPipelineEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetPipelineEXT.html)
- [vkUpdateIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetShaderEXT.html)

## [](#_new_structures)New Structures

- [VkBindIndexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandEXT.html)
- [VkBindVertexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVertexBufferIndirectCommandEXT.html)
- [VkDrawIndirectCountIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCountIndirectCommandEXT.html)
- [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html)
- [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html)
- [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html)
- [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html)
- [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html)
- [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html)
- [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html)
- [VkIndirectCommandsVertexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsVertexBufferTokenEXT.html)
- [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html)
- [VkIndirectExecutionSetPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetPipelineInfoEXT.html)
- [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html)
- [VkIndirectExecutionSetShaderLayoutInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderLayoutInfoEXT.html)
- [VkWriteIndirectExecutionSetPipelineEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteIndirectExecutionSetPipelineEXT.html)
- Extending [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html), [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html):
  
  - [VkGeneratedCommandsPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsPipelineInfoEXT.html)
  - [VkGeneratedCommandsShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsShaderInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)

If [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html) is supported:

- [VkWriteIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteIndirectExecutionSetShaderEXT.html)

## [](#_new_unions)New Unions

- [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html)
- [VkIndirectExecutionSetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoEXT.html)

## [](#_new_enums)New Enums

- [VkIndirectCommandsInputModeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagBitsEXT.html)
- [VkIndirectCommandsLayoutUsageFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagBitsEXT.html)
- [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html)
- [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkIndirectCommandsInputModeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagsEXT.html)
- [VkIndirectCommandsLayoutUsageFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutUsageFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEVICE_GENERATED_COMMANDS_EXTENSION_NAME`
- `VK_EXT_DEVICE_GENERATED_COMMANDS_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_EXT`
  - `VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_EXT`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_PREPROCESS_BUFFER_BIT_EXT`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_EXT`
  - `VK_OBJECT_TYPE_INDIRECT_EXECUTION_SET_EXT`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_EXT`
- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_INDIRECT_BINDABLE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_MEMORY_REQUIREMENTS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_PIPELINE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_SHADER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_TOKEN_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_PIPELINE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_SHADER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_SHADER_LAYOUT_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_WRITE_INDIRECT_EXECUTION_SET_PIPELINE_EXT`
  - `VK_STRUCTURE_TYPE_WRITE_INDIRECT_EXECUTION_SET_SHADER_EXT`

## [](#_example_code)Example Code

TODO

## [](#_version_history)Version History

- Revision 1, 2024-02-23 (Mike Blumenkrantz)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_device_generated_commands)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0