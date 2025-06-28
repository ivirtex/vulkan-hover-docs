# VkIndirectExecutionSetPipelineInfoEXT(3) Manual Page

## Name

VkIndirectExecutionSetPipelineInfoEXT - Struct specifying parameters of a newly created indirect execution set containing only pipelines



## [](#_c_specification)C Specification

The `VkIndirectExecutionSetPipelineInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectExecutionSetPipelineInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         initialPipeline;
    uint32_t           maxPipelineCount;
} VkIndirectExecutionSetPipelineInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `initialPipeline` is the initial pipeline for the set. This pipeline will be automatically added to the set at index `0`.
- `maxPipelineCount` is the maximum number of pipelines stored in the set.

## [](#_description)Description

The characteristics of `initialPipeline` will be used to validate all pipelines added to the set even if they are removed from the set or destroyed.

When an Indirect Execution Set created with pipelines is used, `initialPipeline` constitutes the initial shader state.

Valid Usage

- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11015)VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11015  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesPipelineBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesPipelineBinding` does not contain `VK_SHADER_STAGE_COMPUTE_BIT`, the `VkPipelineBindPoint` of `initialPipeline` **must** not be `VK_PIPELINE_BIND_POINT_COMPUTE`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11016)VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11016  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesPipelineBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesPipelineBinding` does not contain `VK_SHADER_STAGE_FRAGMENT_BIT`, the `VkPipelineBindPoint` of `initialPipeline` **must** not be `VK_PIPELINE_BIND_POINT_GRAPHICS`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11017)VUID-VkIndirectExecutionSetPipelineInfoEXT-supportedIndirectCommandsShaderStagesPipelineBinding-11017  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesPipelineBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesPipelineBinding` does not contain ray tracing stages, the `VkPipelineBindPoint` of `initialPipeline` **must** not be `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-maxPipelineCount-11018)VUID-VkIndirectExecutionSetPipelineInfoEXT-maxPipelineCount-11018  
  `maxPipelineCount` **must** be between `1` and `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT`::`maxIndirectPipelineCount`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-11019)VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-11019  
  `initialPipeline` **must** not use descriptors of type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-11153)VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-11153  
  `initialPipeline` **must** have been created with `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-sType-sType)VUID-VkIndirectExecutionSetPipelineInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_PIPELINE_INFO_EXT`
- [](#VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-parameter)VUID-VkIndirectExecutionSetPipelineInfoEXT-initialPipeline-parameter  
  `initialPipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectExecutionSetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoEXT.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetPipelineInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0