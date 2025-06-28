# VkWriteIndirectExecutionSetPipelineEXT(3) Manual Page

## Name

VkWriteIndirectExecutionSetPipelineEXT - Struct specifying pipeline update information for an indirect execution set



## [](#_c_specification)C Specification

The `VkWriteIndirectExecutionSetPipelineEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkWriteIndirectExecutionSetPipelineEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           index;
    VkPipeline         pipeline;
} VkWriteIndirectExecutionSetPipelineEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `index` is the element of the set to update
- `pipeline` is the pipeline to store in the indirect execution set

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-index-11026)VUID-VkWriteIndirectExecutionSetPipelineEXT-index-11026  
  `index` **must** be less than the value of `VkIndirectExecutionSetPipelineInfoEXT`::`maxPipelineCount` used to create the set
- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-11027)VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-11027  
  `pipeline` **must** have been created with `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`
- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-index-11029)VUID-VkWriteIndirectExecutionSetPipelineEXT-index-11029  
  `index` **must** not be referenced by submitted command buffers
- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-11030)VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-11030  
  The shader stages contained in `pipeline` **must** be supported by [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesPipelineBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesPipelineBinding`

Valid Usage (Implicit)

- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-sType-sType)VUID-VkWriteIndirectExecutionSetPipelineEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_INDIRECT_EXECUTION_SET_PIPELINE_EXT`
- [](#VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-parameter)VUID-VkWriteIndirectExecutionSetPipelineEXT-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkUpdateIndirectExecutionSetPipelineEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetPipelineEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteIndirectExecutionSetPipelineEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0