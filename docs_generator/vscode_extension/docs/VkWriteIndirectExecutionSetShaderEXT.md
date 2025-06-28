# VkWriteIndirectExecutionSetShaderEXT(3) Manual Page

## Name

VkWriteIndirectExecutionSetShaderEXT - Struct specifying shader object update information for an indirect execution set



## [](#_c_specification)C Specification

The `VkWriteIndirectExecutionSetShaderEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands with VK_EXT_shader_object
typedef struct VkWriteIndirectExecutionSetShaderEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           index;
    VkShaderEXT        shader;
} VkWriteIndirectExecutionSetShaderEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `index` is the element of the set to update
- `shader` is the shader to store in the indirect execution set

## [](#_description)Description

Shaders need not be stored in the Indirect Execution Set according to their stage. The only restriction for shader indices within a set is that the value of the index **must** be less than the maximum number of shaders in the set.

Valid Usage

- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-index-11031)VUID-VkWriteIndirectExecutionSetShaderEXT-index-11031  
  `index` **must** be less than `VkIndirectExecutionSetShaderInfoEXT`::`maxShaderCount`
- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-shader-11032)VUID-VkWriteIndirectExecutionSetShaderEXT-shader-11032  
  `shader` **must** have been created with `VK_SHADER_CREATE_INDIRECT_BINDABLE_BIT_EXT`
- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-pInitialShaders-11033)VUID-VkWriteIndirectExecutionSetShaderEXT-pInitialShaders-11033  
  A shader created with the same [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) **must** have been passed in the `VkIndirectExecutionSetShaderInfoEXT`::`pInitialShaders` array
- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-index-11034)VUID-VkWriteIndirectExecutionSetShaderEXT-index-11034  
  `index` **must** not be in use by submitted command buffers

Valid Usage (Implicit)

- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-sType-sType)VUID-VkWriteIndirectExecutionSetShaderEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_INDIRECT_EXECUTION_SET_SHADER_EXT`
- [](#VUID-VkWriteIndirectExecutionSetShaderEXT-shader-parameter)VUID-VkWriteIndirectExecutionSetShaderEXT-shader-parameter  
  `shader` **must** be a valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handle

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkUpdateIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetShaderEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteIndirectExecutionSetShaderEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0