# VkIndirectCommandsTokenDataEXT(3) Manual Page

## Name

VkIndirectCommandsTokenDataEXT - Union specifying the token-specific details of an indirect command layout token



## [](#_c_specification)C Specification

The `VkIndirectCommandsTokenDataEXT` structure provides token-specific details used to generate the indirect execution layout.

```c++
// Provided by VK_EXT_device_generated_commands
typedef union VkIndirectCommandsTokenDataEXT {
    const VkIndirectCommandsPushConstantTokenEXT*    pPushConstant;
    const VkIndirectCommandsVertexBufferTokenEXT*    pVertexBuffer;
    const VkIndirectCommandsIndexBufferTokenEXT*     pIndexBuffer;
    const VkIndirectCommandsExecutionSetTokenEXT*    pExecutionSet;
} VkIndirectCommandsTokenDataEXT;
```

## [](#_members)Members

- `pPushConstant` is a pointer to a [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html) structure needed for `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` and `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT` tokens
- `pVertexBuffer` is a pointer to a [VkIndirectCommandsVertexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsVertexBufferTokenEXT.html) structure needed for `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT` tokens
- `pIndexBuffer` is a pointer to a [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html) structure needed for `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT` tokens
- `pExecutionSet` is a pointer to a [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html) structure needed for `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` tokens

## [](#_description)Description

The appropriate member of the union **must** be set for each token.

The following code provides detailed information on how an individual sequence is processed. For valid usage, all restrictions from the regular commands apply.

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html), [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html), [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html), [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html), [VkIndirectCommandsVertexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsVertexBufferTokenEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsTokenDataEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0