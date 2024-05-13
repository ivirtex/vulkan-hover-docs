# VkCommandBufferBeginInfo(3) Manual Page

## Name

VkCommandBufferBeginInfo - Structure specifying a command buffer begin
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCommandBufferBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkCommandBufferBeginInfo {
    VkStructureType                          sType;
    const void*                              pNext;
    VkCommandBufferUsageFlags                flags;
    const VkCommandBufferInheritanceInfo*    pInheritanceInfo;
} VkCommandBufferBeginInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkCommandBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html)
  specifying usage behavior for the command buffer.

- `pInheritanceInfo` is a pointer to a
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  structure, used if `commandBuffer` is a secondary command buffer. If
  this is a primary command buffer, then this value is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCommandBufferBeginInfo-flags-09123"
  id="VUID-VkCommandBufferBeginInfo-flags-09123"></a>
  VUID-VkCommandBufferBeginInfo-flags-09123  
  If `flags` contains
  `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) that `commandBuffer` was allocated
  from **must** support graphics operations

- <a href="#VUID-VkCommandBufferBeginInfo-flags-00055"
  id="VUID-VkCommandBufferBeginInfo-flags-00055"></a>
  VUID-VkCommandBufferBeginInfo-flags-00055  
  If `flags` contains
  `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the `framebuffer`
  member of `pInheritanceInfo` **must** be either
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), or a valid `VkFramebuffer` that
  is compatible with the `renderPass` member of `pInheritanceInfo`

- <a href="#VUID-VkCommandBufferBeginInfo-flags-09240"
  id="VUID-VkCommandBufferBeginInfo-flags-09240"></a>
  VUID-VkCommandBufferBeginInfo-flags-09240  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`
  and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRendering"
  target="_blank" rel="noopener"><code>dynamicRendering</code></a>
  feature is not enabled, the `renderPass` member of `pInheritanceInfo`
  **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkCommandBufferBeginInfo-flags-06002"
  id="VUID-VkCommandBufferBeginInfo-flags-06002"></a>
  VUID-VkCommandBufferBeginInfo-flags-06002  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`
  and the `renderPass` member of `pInheritanceInfo` is
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext` chain of
  `pInheritanceInfo` **must** include a
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure

- <a href="#VUID-VkCommandBufferBeginInfo-flags-06003"
  id="VUID-VkCommandBufferBeginInfo-flags-06003"></a>
  VUID-VkCommandBufferBeginInfo-flags-06003  
  If `flags` contains
  `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`, the `renderPass`
  member of `pInheritanceInfo` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  and the `pNext` chain of `pInheritanceInfo` includes a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, the `colorAttachmentCount` member of that structure
  **must** be equal to the value of
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`colorAttachmentCount`

- <a href="#VUID-VkCommandBufferBeginInfo-flags-06000"
  id="VUID-VkCommandBufferBeginInfo-flags-06000"></a>
  VUID-VkCommandBufferBeginInfo-flags-06000  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`
  and the `renderPass` member of `pInheritanceInfo` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `renderPass` member of
  `pInheritanceInfo` **must** be a valid `VkRenderPass`

- <a href="#VUID-VkCommandBufferBeginInfo-flags-06001"
  id="VUID-VkCommandBufferBeginInfo-flags-06001"></a>
  VUID-VkCommandBufferBeginInfo-flags-06001  
  If `flags` contains `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`
  and the `renderPass` member of `pInheritanceInfo` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `subpass` member of
  `pInheritanceInfo` **must** be a valid subpass index within the
  `renderPass` member of `pInheritanceInfo`

Valid Usage (Implicit)

- <a href="#VUID-VkCommandBufferBeginInfo-sType-sType"
  id="VUID-VkCommandBufferBeginInfo-sType-sType"></a>
  VUID-VkCommandBufferBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO`

- <a href="#VUID-VkCommandBufferBeginInfo-pNext-pNext"
  id="VUID-VkCommandBufferBeginInfo-pNext-pNext"></a>
  VUID-VkCommandBufferBeginInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html)

- <a href="#VUID-VkCommandBufferBeginInfo-sType-unique"
  id="VUID-VkCommandBufferBeginInfo-sType-unique"></a>
  VUID-VkCommandBufferBeginInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkCommandBufferBeginInfo-flags-parameter"
  id="VUID-VkCommandBufferBeginInfo-flags-parameter"></a>
  VUID-VkCommandBufferBeginInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkCommandBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlagBits.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
[VkCommandBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
