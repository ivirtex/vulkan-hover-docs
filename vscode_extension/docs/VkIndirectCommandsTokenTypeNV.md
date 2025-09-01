# VkIndirectCommandsTokenTypeNV(3) Manual Page

## Name

VkIndirectCommandsTokenTypeNV - Enum specifying token commands



## [](#_c_specification)C Specification

Possible values of those elements of the [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html)::`pTokens` array specifying command tokens (other elements of the array specify command parameters) are:

```c++
// Provided by VK_NV_device_generated_commands
typedef enum VkIndirectCommandsTokenTypeNV {
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV = 0,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV = 1,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV = 2,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV = 3,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV = 4,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV = 5,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV = 6,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV = 7,
  // Provided by VK_EXT_mesh_shader with VK_NV_device_generated_commands
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV = 1000328000,
  // Provided by VK_NV_device_generated_commands_compute
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV = 1000428003,
  // Provided by VK_NV_device_generated_commands_compute
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV = 1000428004,
} VkIndirectCommandsTokenTypeNV;
```

## [](#_description)Description

Table 1. Supported Indirect Command Tokens   Token type Equivalent command

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV`

[vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipelineShaderGroupNV.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV`

\-

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV`

[vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV`

[vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV`

[vkCmdPushConstants](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV`

[vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirect.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV`

[vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirect.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV`

[vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectNV.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV`

[vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectEXT.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`

[vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipeline.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV`

[vkCmdDispatchIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchIndirect.html)

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsTokenTypeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0