# VkGeneratedCommandsMemoryRequirementsInfoNV(3) Manual Page

## Name

VkGeneratedCommandsMemoryRequirementsInfoNV - Structure specifying parameters for the reservation of preprocess buffer space



## [](#_c_specification)C Specification

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkGeneratedCommandsMemoryRequirementsInfoNV {
    VkStructureType               sType;
    const void*                   pNext;
    VkPipelineBindPoint           pipelineBindPoint;
    VkPipeline                    pipeline;
    VkIndirectCommandsLayoutNV    indirectCommandsLayout;
    uint32_t                      maxSequencesCount;
} VkGeneratedCommandsMemoryRequirementsInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineBindPoint` is the [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) of the `pipeline` that this buffer memory is intended to be used with during the execution.
- `pipeline` is the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) that this buffer memory is intended to be used with during the execution.
- `indirectCommandsLayout` is the [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html) that this buffer memory is intended to be used with.
- `maxSequencesCount` is the maximum number of sequences that this buffer memory in combination with the other state provided **can** be used with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-maxSequencesCount-02907)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-maxSequencesCount-02907  
  `maxSequencesCount` **must** be less or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV.html)::`maxIndirectSequenceCount`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09075)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09075  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_GRAPHICS`, then `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09076)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09076  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`, and the `indirectCommandsLayout` was not created with a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token, then the `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09077)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09077  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`, and the `indirectCommandsLayout` contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token, then the `pipeline` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-sType-sType)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_MEMORY_REQUIREMENTS_INFO_NV`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pNext-pNext)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-parameter)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipeline-parameter)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipeline-parameter  
  If `pipeline` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-indirectCommandsLayout-parameter)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-indirectCommandsLayout-parameter  
  `indirectCommandsLayout` **must** be a valid [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-commonparent)VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-commonparent  
  Both of `indirectCommandsLayout`, and `pipeline` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetGeneratedCommandsMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeneratedCommandsMemoryRequirementsInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0