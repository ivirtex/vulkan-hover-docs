# VkIndirectCommandsPushConstantTokenEXT(3) Manual Page

## Name

VkIndirectCommandsPushConstantTokenEXT - Structure specifying layout token info for a single push constant command token



## [](#_c_specification)C Specification

The `VkIndirectCommandsPushConstantTokenEXT` structure specifies the layout token info for `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` and `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` tokens.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsPushConstantTokenEXT {
    VkPushConstantRange    updateRange;
} VkIndirectCommandsPushConstantTokenEXT;
```

## [](#_members)Members

- `updateRange` is the push constant range that will be updated by the token.

## [](#_description)Description

The `stageFlags` member of `updateRange` is ignored.

Valid Usage

- [](#VUID-VkIndirectCommandsPushConstantTokenEXT-updateRange-11132)VUID-VkIndirectCommandsPushConstantTokenEXT-updateRange-11132  
  `updateRange` **must** be contained within the push constant info used by [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html)
- [](#VUID-VkIndirectCommandsPushConstantTokenEXT-size-11133)VUID-VkIndirectCommandsPushConstantTokenEXT-size-11133  
  If the token type is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`, the `size` member of `updateRange` **must** be 4

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsPushConstantTokenEXT-updateRange-parameter)VUID-VkIndirectCommandsPushConstantTokenEXT-updateRange-parameter  
  `updateRange` **must** be a valid [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) structure

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html), [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsPushConstantTokenEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0