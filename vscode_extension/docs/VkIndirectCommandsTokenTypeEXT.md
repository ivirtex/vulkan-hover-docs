# VkIndirectCommandsTokenTypeEXT(3) Manual Page

## Name

VkIndirectCommandsTokenTypeEXT - Enum specifying token commands



## [](#_c_specification)C Specification

Possible values of those elements of the [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html)::`pTokens` array specifying command tokens (other elements of the array specify command parameters) are:

```c++
// Provided by VK_EXT_device_generated_commands
typedef enum VkIndirectCommandsTokenTypeEXT {
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT = 0,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT = 1,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT = 2,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT = 3,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT = 4,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_EXT = 5,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_EXT = 6,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_COUNT_EXT = 7,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_COUNT_EXT = 8,
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_EXT = 9,
  // Provided by VK_EXT_device_generated_commands with VK_NV_mesh_shader
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT = 1000202002,
  // Provided by VK_EXT_device_generated_commands with VK_NV_mesh_shader
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT = 1000202003,
  // Provided by VK_EXT_device_generated_commands with VK_EXT_mesh_shader
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT = 1000328000,
  // Provided by VK_EXT_device_generated_commands with VK_EXT_mesh_shader
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT = 1000328001,
  // Provided by VK_KHR_ray_tracing_maintenance1 with VK_EXT_device_generated_commands
    VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT = 1000386004,
} VkIndirectCommandsTokenTypeEXT;
```

## [](#_description)Description

Table 1. Supported Indirect Command Tokens   **Common Tokens** **Command Data**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT`

`u32[]` array of indices into the indirect execution set

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT`

`u32[]` raw data

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`

`u32` placeholder data (not accessed by shader)

**Compute Tokens**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_EXT`

[VkDispatchIndirectCommand](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchIndirectCommand.html)

**Ray Tracing Tokens**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`

[VkTraceRaysIndirectCommand2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTraceRaysIndirectCommand2KHR.html)

**Graphics State Tokens**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT`

[VkBindIndexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandEXT.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT`

[VkBindVertexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVertexBufferIndirectCommandEXT.html)

**Graphics Draw Tokens**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_EXT`

[VkDrawIndexedIndirectCommand](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndexedIndirectCommand.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_EXT`

[VkDrawIndirectCommand](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCommand.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT`

[VkDrawMeshTasksIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawMeshTasksIndirectCommandEXT.html)

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT`

[VkDrawMeshTasksIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawMeshTasksIndirectCommandNV.html)

**Graphics Draw Count Tokens**

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_COUNT_EXT`

[VkDrawIndirectCountIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCountIndirectCommandEXT.html) with VkDrawIndexedIndirectCommand

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_COUNT_EXT`

[VkDrawIndirectCountIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCountIndirectCommandEXT.html) with VkDrawIndirectCommand

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`

[VkDrawIndirectCountIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCountIndirectCommandEXT.html) with VkDrawMeshTasksIndirectCommandEXT

`VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`

[VkDrawIndirectCountIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawIndirectCountIndirectCommandEXT.html) with VkDrawMeshTasksIndirectCommandNV

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsTokenTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0