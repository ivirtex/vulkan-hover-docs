# VkSubpassBeginInfo(3) Manual Page

## Name

VkSubpassBeginInfo - Structure specifying subpass begin information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSubpassBeginInfo {
    VkStructureType      sType;
    const void*          pNext;
    VkSubpassContents    contents;
} VkSubpassBeginInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassBeginInfo VkSubpassBeginInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `contents` is a [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html) value
  specifying how the commands in the next subpass will be provided.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSubpassBeginInfo-contents-09382"
  id="VUID-VkSubpassBeginInfo-contents-09382"></a>
  VUID-VkSubpassBeginInfo-contents-09382  
  If `contents` is
  `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT`, then
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassBeginInfo-sType-sType"
  id="VUID-VkSubpassBeginInfo-sType-sType"></a>
  VUID-VkSubpassBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO`

- <a href="#VUID-VkSubpassBeginInfo-pNext-pNext"
  id="VUID-VkSubpassBeginInfo-pNext-pNext"></a>
  VUID-VkSubpassBeginInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSubpassBeginInfo-contents-parameter"
  id="VUID-VkSubpassBeginInfo-contents-parameter"></a>
  VUID-VkSubpassBeginInfo-contents-parameter  
  `contents` **must** be a valid
  [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html),
[vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html),
[vkCmdBeginRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2KHR.html),
[vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html),
[vkCmdNextSubpass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
