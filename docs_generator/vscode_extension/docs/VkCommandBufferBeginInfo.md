# VkCommandBufferBeginInfo(3) Manual Page

## Name

VkCommandBufferBeginInfo - Structure specifying a command buffer begin operation



## [](#_c_specification)C Specification

The `VkCommandBufferBeginInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkCommandBufferBeginInfo {
    VkStructureType                          sType;
    const void*                              pNext;
    VkCommandBufferUsageFlags                flags;
    const VkCommandBufferInheritanceInfo*    pInheritanceInfo;
} VkCommandBufferBeginInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkCommandBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferUsageFlagBits.html) specifying usage behavior for the command buffer.
- `pInheritanceInfo` is a pointer to a [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) structure, used if `commandBuffer` is a secondary command buffer. If this is a primary command buffer, then this value is ignored.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCommandBufferBeginInfo-flags-09123)VUID-VkCommandBufferBeginInfo-flags-09123  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-VkCommandBufferBeginInfo-flags-00055)VUID-VkCommandBufferBeginInfo-flags-00055  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the `framebuffer` member of `pInheritanceInfo` **must** be either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or a valid `VkFramebuffer` that is compatible with the `renderPass` member of `pInheritanceInfo`
- [](#VUID-VkCommandBufferBeginInfo-flags-09240)VUID-VkCommandBufferBeginInfo-flags-09240  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` and the [`dynamicRendering`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRendering) feature is not enabled, the `renderPass` member of `pInheritanceInfo` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkCommandBufferBeginInfo-flags-06002)VUID-VkCommandBufferBeginInfo-flags-06002  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` and the `renderPass` member of `pInheritanceInfo` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain of `pInheritanceInfo` **must** include a [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html) structure
- [](#VUID-VkCommandBufferBeginInfo-flags-06003)VUID-VkCommandBufferBeginInfo-flags-06003  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the `renderPass` member of `pInheritanceInfo` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and the `pNext` chain of `pInheritanceInfo` includes a [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleCountInfoAMD.html) or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleCountInfoNV.html) structure, the `colorAttachmentCount` member of that structure **must** be equal to the value of [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`colorAttachmentCount`
- [](#VUID-VkCommandBufferBeginInfo-flags-06000)VUID-VkCommandBufferBeginInfo-flags-06000  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` and the `renderPass` member of `pInheritanceInfo` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `renderPass` member of `pInheritanceInfo` **must** be a valid `VkRenderPass`
- [](#VUID-VkCommandBufferBeginInfo-flags-06001)VUID-VkCommandBufferBeginInfo-flags-06001  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` and the `renderPass` member of `pInheritanceInfo` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `subpass` member of `pInheritanceInfo` **must** be a valid subpass index within the `renderPass` member of `pInheritanceInfo`
- [](#VUID-VkCommandBufferBeginInfo-flags-10617)VUID-VkCommandBufferBeginInfo-flags-10617  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` , the `renderPass` member of `pInheritanceInfo` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and `renderPass` was created with [tile shading enabled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading), `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` **must** be included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`
- [](#VUID-VkCommandBufferBeginInfo-flags-10618)VUID-VkCommandBufferBeginInfo-flags-10618  
  If `flags` does not contain `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` , the `renderPass` member of `pInheritanceInfo` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or `renderPass` was not created with tile shading enabled, `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` **must** not be included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`
- [](#VUID-VkCommandBufferBeginInfo-flags-10619)VUID-VkCommandBufferBeginInfo-flags-10619  
  If `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` is included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`, [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`tileApronSize` **must** be equal to the `tileApronSize` used to create `renderPass`

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferBeginInfo-sType-sType)VUID-VkCommandBufferBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO`
- [](#VUID-VkCommandBufferBeginInfo-pNext-pNext)VUID-VkCommandBufferBeginInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupCommandBufferBeginInfo.html)
- [](#VUID-VkCommandBufferBeginInfo-sType-unique)VUID-VkCommandBufferBeginInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkCommandBufferBeginInfo-flags-parameter)VUID-VkCommandBufferBeginInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkCommandBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferUsageFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), [VkCommandBufferUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBeginCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferBeginInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0