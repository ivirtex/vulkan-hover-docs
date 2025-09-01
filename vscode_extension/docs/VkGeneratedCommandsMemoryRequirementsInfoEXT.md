# VkGeneratedCommandsMemoryRequirementsInfoEXT(3) Manual Page

## Name

VkGeneratedCommandsMemoryRequirementsInfoEXT - Structure specifying parameters for the reservation of preprocess buffer space



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkGeneratedCommandsMemoryRequirementsInfoEXT {
    VkStructureType                sType;
    const void*                    pNext;
    VkIndirectExecutionSetEXT      indirectExecutionSet;
    VkIndirectCommandsLayoutEXT    indirectCommandsLayout;
    uint32_t                       maxSequenceCount;
    uint32_t                       maxDrawCount;
} VkGeneratedCommandsMemoryRequirementsInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `indirectExecutionSet` is the indirect execution set to be used for binding shaders.
- `indirectCommandsLayout` is the [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) that this buffer memory is intended to be used with.
- `maxSequenceCount` is the maximum number of sequences that this buffer memory can be used with.
- `maxDrawCount` is the maximum number of indirect draws that can be executed by any COUNT-type multi-draw indirect tokens. The draw count in the indirect buffer is clamped to this value for these token types.

## [](#_description)Description

If the action command token for the layout is not a COUNT-type multi-draw indirect token, `maxDrawCount` is ignored.

Valid Usage

- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-maxSequencesCount-11009)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-maxSequencesCount-11009  
  `maxSequencesCount` **must** be less or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectSequenceCount`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11010)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11010  
  If `indirectCommandsLayout` was created with a token sequence that contained the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, `indirectExecutionSet` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11151)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11151  
  If `indirectCommandsLayout` was created with a token sequence that contained the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, the shader stages used to create the initial shader state of `indirectExecutionSet` **must** equal the [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html)::`shaderStages` used to create `indirectCommandsLayout`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11011)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-11011  
  If `indirectCommandsLayout` was not created with a token sequence that contained the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, `indirectExecutionSet` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-maxDrawCount-11146)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-maxDrawCount-11146  
  When not ignored, `maxDrawCount` × `maxSequenceCount` **must** be less than 2^24
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectExecutionSet-11012)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectExecutionSet-11012  
  If `indirectExecutionSet` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), either a [VkGeneratedCommandsPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsPipelineInfoEXT.html) or a [VkGeneratedCommandsShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsShaderInfoEXT.html) **must** be included in the `pNext` chain

Valid Usage (Implicit)

- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-sType-sType)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_MEMORY_REQUIREMENTS_INFO_EXT`
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-pNext-pNext)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkGeneratedCommandsPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsPipelineInfoEXT.html) or [VkGeneratedCommandsShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsShaderInfoEXT.html)
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-sType-unique)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectExecutionSet-parameter)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectExecutionSet-parameter  
  If `indirectExecutionSet` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indirectExecutionSet` **must** be a valid [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-parameter)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-indirectCommandsLayout-parameter  
  `indirectCommandsLayout` **must** be a valid [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) handle
- [](#VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-commonparent)VUID-VkGeneratedCommandsMemoryRequirementsInfoEXT-commonparent  
  Both of `indirectCommandsLayout`, and `indirectExecutionSet` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetGeneratedCommandsMemoryRequirementsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetGeneratedCommandsMemoryRequirementsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeneratedCommandsMemoryRequirementsInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0