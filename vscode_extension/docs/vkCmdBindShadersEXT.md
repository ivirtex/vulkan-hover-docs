# vkCmdBindShadersEXT(3) Manual Page

## Name

vkCmdBindShadersEXT - Bind shader objects to a command buffer



## [](#_c_specification)C Specification

Once shader objects have been created, they **can** be bound to the command buffer using the command:

```c++
// Provided by VK_EXT_shader_object
void vkCmdBindShadersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    stageCount,
    const VkShaderStageFlagBits*                pStages,
    const VkShaderEXT*                          pShaders);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the shader object will be bound to.
- `stageCount` is the length of the `pStages` and `pShaders` arrays.
- `pStages` is a pointer to an array of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values specifying one stage per array index that is affected by the corresponding value in the `pShaders` array.
- `pShaders` is a pointer to an array of `VkShaderEXT` handles and/or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) values describing the shader binding operations to be performed on each stage in `pStages`.

## [](#_description)Description

When binding linked shaders, an application **may** bind them in any combination of one or more calls to `vkCmdBindShadersEXT` (i.e., shaders that were created linked together do not need to be bound in the same `vkCmdBindShadersEXT` call).

Any shader object bound to a particular stage **may** be unbound by setting its value in `pShaders` to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). If `pShaders` is `NULL`, `vkCmdBindShadersEXT` behaves as if `pShaders` was an array of `stageCount` [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) values (i.e., any shaders bound to the stages specified in `pStages` are unbound).

Valid Usage

- [](#VUID-vkCmdBindShadersEXT-None-08462)VUID-vkCmdBindShadersEXT-None-08462  
  The [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature **must** be enabled
- [](#VUID-vkCmdBindShadersEXT-pStages-08463)VUID-vkCmdBindShadersEXT-pStages-08463  
  Every element of `pStages` **must** be unique
- [](#VUID-vkCmdBindShadersEXT-pStages-08464)VUID-vkCmdBindShadersEXT-pStages-08464  
  `pStages` **must** not contain `VK_SHADER_STAGE_ALL_GRAPHICS` or `VK_SHADER_STAGE_ALL`
- [](#VUID-vkCmdBindShadersEXT-pStages-08465)VUID-vkCmdBindShadersEXT-pStages-08465  
  `pStages` **must** not contain `VK_SHADER_STAGE_RAYGEN_BIT_KHR`, `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`, `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`, `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`, or `VK_SHADER_STAGE_CALLABLE_BIT_KHR`
- [](#VUID-vkCmdBindShadersEXT-pStages-08467)VUID-vkCmdBindShadersEXT-pStages-08467  
  `pStages` **must** not contain `VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI`
- [](#VUID-vkCmdBindShadersEXT-pStages-08468)VUID-vkCmdBindShadersEXT-pStages-08468  
  `pStages` **must** not contain `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`
- [](#VUID-vkCmdBindShadersEXT-pShaders-08469)VUID-vkCmdBindShadersEXT-pShaders-08469  
  For each element of `pStages`, if `pShaders` is not `NULL`, and the element of the `pShaders` array with the same index is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have been created with a `stage` equal to the corresponding element of `pStages`
- [](#VUID-vkCmdBindShadersEXT-pShaders-08470)VUID-vkCmdBindShadersEXT-pShaders-08470  
  If `pStages` contains both `VK_SHADER_STAGE_TASK_BIT_EXT` and `VK_SHADER_STAGE_VERTEX_BIT`, and `pShaders` is not `NULL`, and the same index in `pShaders` as `VK_SHADER_STAGE_TASK_BIT_EXT` in `pStages` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the same index in `pShaders` as `VK_SHADER_STAGE_VERTEX_BIT` in `pStages` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBindShadersEXT-pShaders-08471)VUID-vkCmdBindShadersEXT-pShaders-08471  
  If `pStages` contains both `VK_SHADER_STAGE_MESH_BIT_EXT` and `VK_SHADER_STAGE_VERTEX_BIT`, and `pShaders` is not `NULL`, and the same index in `pShaders` as `VK_SHADER_STAGE_MESH_BIT_EXT` in `pStages` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the same index in `pShaders` as `VK_SHADER_STAGE_VERTEX_BIT` in `pStages` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBindShadersEXT-pShaders-08476)VUID-vkCmdBindShadersEXT-pShaders-08476  
  If `pStages` contains `VK_SHADER_STAGE_COMPUTE_BIT`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBindShadersEXT-pShaders-08477)VUID-vkCmdBindShadersEXT-pShaders-08477  
  If `pStages` contains `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `VK_SHADER_STAGE_GEOMETRY_BIT`, or `VK_SHADER_STAGE_FRAGMENT_BIT`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindShadersEXT-pShaders-08478)VUID-vkCmdBindShadersEXT-pShaders-08478  
  If `pStages` contains `VK_SHADER_STAGE_MESH_BIT_EXT` or `VK_SHADER_STAGE_TASK_BIT_EXT`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations

Valid Usage (Implicit)

- [](#VUID-vkCmdBindShadersEXT-commandBuffer-parameter)VUID-vkCmdBindShadersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindShadersEXT-pStages-parameter)VUID-vkCmdBindShadersEXT-pStages-parameter  
  `pStages` **must** be a valid pointer to an array of `stageCount` valid [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-vkCmdBindShadersEXT-pShaders-parameter)VUID-vkCmdBindShadersEXT-pShaders-parameter  
  If `pShaders` is not `NULL`, `pShaders` **must** be a valid pointer to an array of `stageCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handles
- [](#VUID-vkCmdBindShadersEXT-commandBuffer-recording)VUID-vkCmdBindShadersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindShadersEXT-commandBuffer-cmdpool)VUID-vkCmdBindShadersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdBindShadersEXT-videocoding)VUID-vkCmdBindShadersEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindShadersEXT-stageCount-arraylength)VUID-vkCmdBindShadersEXT-stageCount-arraylength  
  `stageCount` **must** be greater than `0`
- [](#VUID-vkCmdBindShadersEXT-commonparent)VUID-vkCmdBindShadersEXT-commonparent  
  Both of `commandBuffer`, and the elements of `pShaders` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute

State

Conditional Rendering

vkCmdBindShadersEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html), [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindShadersEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0