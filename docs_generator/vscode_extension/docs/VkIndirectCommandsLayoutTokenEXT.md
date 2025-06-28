# VkIndirectCommandsLayoutTokenEXT(3) Manual Page

## Name

VkIndirectCommandsLayoutTokenEXT - Struct specifying the details of an indirect command layout token



## [](#_c_specification)C Specification

The `VkIndirectCommandsLayoutTokenEXT` structure specifies details to the function arguments that need to be known at layout creation time:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsLayoutTokenEXT {
    VkStructureType                   sType;
    const void*                       pNext;
    VkIndirectCommandsTokenTypeEXT    type;
    VkIndirectCommandsTokenDataEXT    data;
    uint32_t                          offset;
} VkIndirectCommandsLayoutTokenEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` specifies the [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html) for `data`.
- `data` specifies a [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html) containing token-specific details for command execution. It is ignored if `type` does not match any member of the [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html) union.
- `offset` is the relative byte offset for the token within one sequence of the indirect buffer. The data stored at that offset is the command data for the token, e.g. `VkDispatchIndirectCommand`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectCommandsLayoutTokenEXT-offset-11124)VUID-VkIndirectCommandsLayoutTokenEXT-offset-11124  
  `offset` **must** be less than or equal to [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`maxIndirectCommandsTokenOffset`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-offset-11125)VUID-VkIndirectCommandsLayoutTokenEXT-offset-11125  
  `offset` **must** be aligned to `4`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-meshShader-11126)VUID-VkIndirectCommandsLayoutTokenEXT-meshShader-11126  
  If [`meshShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-meshShader) or [`taskShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-taskShader) are not enabled, `type` **must** not be `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT` `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`, `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-rayTracingMaintenance1-11128)VUID-VkIndirectCommandsLayoutTokenEXT-rayTracingMaintenance1-11128  
  If the [`rayTracingMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTracingMaintenance1) feature is not enabled, `type` **must** not be `VK_INDIRECT_COMMANDS_TOKEN_TYPE_TRACE_RAYS2_EXT`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11129)VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11129  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-deviceGeneratedCommandsMultiDrawIndirectCount)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`deviceGeneratedCommandsMultiDrawIndirectCount` is not supported, `type` **must** not be `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_COUNT_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_COUNT_EXT`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11130)VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11130  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-deviceGeneratedCommandsMultiDrawIndirectCount)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`deviceGeneratedCommandsMultiDrawIndirectCount` is not supported, `type` **must** not be `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11131)VUID-VkIndirectCommandsLayoutTokenEXT-deviceGeneratedCommandsMultiDrawIndirectCount-11131  
  If [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-deviceGeneratedCommandsMultiDrawIndirectCount)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`deviceGeneratedCommandsMultiDrawIndirectCount` is not supported, `type` **must** not be `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsLayoutTokenEXT-sType-sType)VUID-VkIndirectCommandsLayoutTokenEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_TOKEN_EXT`
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-type-parameter)VUID-VkIndirectCommandsLayoutTokenEXT-type-parameter  
  `type` **must** be a valid [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html) value
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-pPushConstant-parameter)VUID-VkIndirectCommandsLayoutTokenEXT-pPushConstant-parameter  
  If `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_EXT` or `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SEQUENCE_INDEX_EXT`, the `pPushConstant` member of `data` **must** be a valid pointer to a valid [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html) structure
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-pVertexBuffer-parameter)VUID-VkIndirectCommandsLayoutTokenEXT-pVertexBuffer-parameter  
  If `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT`, the `pVertexBuffer` member of `data` **must** be a valid pointer to a valid [VkIndirectCommandsVertexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsVertexBufferTokenEXT.html) structure
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-pIndexBuffer-parameter)VUID-VkIndirectCommandsLayoutTokenEXT-pIndexBuffer-parameter  
  If `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT`, the `pIndexBuffer` member of `data` **must** be a valid pointer to a valid [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html) structure
- [](#VUID-VkIndirectCommandsLayoutTokenEXT-pExecutionSet-parameter)VUID-VkIndirectCommandsLayoutTokenEXT-pExecutionSet-parameter  
  If `type` is `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT`, the `pExecutionSet` member of `data` **must** be a valid pointer to a valid [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html), [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html), [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsLayoutTokenEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0