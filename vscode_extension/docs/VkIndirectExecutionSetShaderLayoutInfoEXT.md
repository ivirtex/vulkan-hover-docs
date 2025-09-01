# VkIndirectExecutionSetShaderLayoutInfoEXT(3) Manual Page

## Name

VkIndirectExecutionSetShaderLayoutInfoEXT - Struct specifying descriptor layout parameters of a newly created indirect execution set containing only shader objects



## [](#_c_specification)C Specification

The `VkIndirectExecutionSetShaderLayoutInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectExecutionSetShaderLayoutInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        setLayoutCount;
    const VkDescriptorSetLayout*    pSetLayouts;
} VkIndirectExecutionSetShaderLayoutInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `setLayoutCount` is the number of members in the `pSetLayouts` array
- `pSetLayouts` is a pointer to an array containing [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) objects used by the shader stage. The implementation **must** not access these objects outside of the duration of the command this structure is passed to.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-pSetLayouts-11024)VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-pSetLayouts-11024  
  All members of `pSetLayouts` **must** not contain descriptors of type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

Valid Usage (Implicit)

- [](#VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-sType-sType)VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_SHADER_LAYOUT_INFO_EXT`
- [](#VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-pSetLayouts-parameter)VUID-VkIndirectExecutionSetShaderLayoutInfoEXT-pSetLayouts-parameter  
  If `setLayoutCount` is not `0`, `pSetLayouts` **must** be a valid pointer to an array of `setLayoutCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handles

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetShaderLayoutInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0