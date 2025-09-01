# VkGeneratedCommandsPipelineInfoEXT(3) Manual Page

## Name

VkGeneratedCommandsPipelineInfoEXT - Structure specifying a pipeline for use with indirect command preprocessing



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkGeneratedCommandsPipelineInfoEXT {
    VkStructureType    sType;
    void*              pNext;
    VkPipeline         pipeline;
} VkGeneratedCommandsPipelineInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipeline` is a valid pipeline object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkGeneratedCommandsPipelineInfoEXT-sType-sType)VUID-VkGeneratedCommandsPipelineInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_PIPELINE_INFO_EXT`
- [](#VUID-VkGeneratedCommandsPipelineInfoEXT-pipeline-parameter)VUID-VkGeneratedCommandsPipelineInfoEXT-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeneratedCommandsPipelineInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0