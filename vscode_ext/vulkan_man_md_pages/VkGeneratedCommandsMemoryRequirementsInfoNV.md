# VkGeneratedCommandsMemoryRequirementsInfoNV(3) Manual Page

## Name

VkGeneratedCommandsMemoryRequirementsInfoNV - Structure specifying
parameters for the reservation of preprocess buffer space



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pipelineBindPoint` is the
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) of the `pipeline` that
  this buffer memory is intended to be used with during the execution.

- `pipeline` is the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) that this buffer
  memory is intended to be used with during the execution.

- `indirectCommandsLayout` is the
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) that
  this buffer memory is intended to be used with.

- `maxSequencesCount` is the maximum number of sequences that this
  buffer memory in combination with the other state provided **can** be
  used with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-maxSequencesCount-02907"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-maxSequencesCount-02907"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-maxSequencesCount-02907  
  `maxSequencesCount` **must** be less or equal to
  [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV.html)::`maxIndirectSequenceCount`

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09075"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09075"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09075  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_GRAPHICS`,
  then `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handle

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09076"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09076"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09076  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  and the `indirectCommandsLayout` was not created with a
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token, then the
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09077"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09077"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-09077  
  If `pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`,
  and the `indirectCommandsLayout` contains a
  `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token, then the
  `pipeline` **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- <a href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-sType-sType"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-sType-sType"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_MEMORY_REQUIREMENTS_INFO_NV`

- <a href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pNext-pNext"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pNext-pNext"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-parameter"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-parameter"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipeline-parameter"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipeline-parameter"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-pipeline-parameter  
  If `pipeline` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pipeline`
  **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-indirectCommandsLayout-parameter"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-indirectCommandsLayout-parameter"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-indirectCommandsLayout-parameter  
  `indirectCommandsLayout` **must** be a valid
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) handle

- <a href="#VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-commonparent"
  id="VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-commonparent"></a>
  VUID-VkGeneratedCommandsMemoryRequirementsInfoNV-commonparent  
  Both of `indirectCommandsLayout`, and `pipeline` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetGeneratedCommandsMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetGeneratedCommandsMemoryRequirementsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeneratedCommandsMemoryRequirementsInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
