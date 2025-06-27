# VkIndirectCommandsTokenTypeNV(3) Manual Page

## Name

VkIndirectCommandsTokenTypeNV - Enum specifying token commands



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of those elements of the
[VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html)::`pTokens`
array specifying command tokens (other elements of the array specify
command parameters) are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

| Token type | Equivalent command |
|----|----|
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV` | [vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipelineShaderGroupNV.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_STATE_FLAGS_NV` | \- |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV` | [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV` | [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NV` | [vkCmdPushConstants](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NV` | [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NV` | [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_TASKS_NV` | [vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectNV.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV` | [vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectEXT.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` | [vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html) |
| `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV` | [vkCmdDispatchIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchIndirect.html) |

Table 1. Supported indirect command tokens

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsTokenTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
