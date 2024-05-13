# VkCommandPool(3) Manual Page

## Name

VkCommandPool - Opaque handle to a command pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

Command pools are opaque objects that command buffer memory is allocated
from, and which allow the implementation to amortize the cost of
resource creation across multiple command buffers. Command pools are
externally synchronized, meaning that a command pool **must** not be
used concurrently in multiple threads. That includes use via recording
commands on any command buffers allocated from the pool, as well as
operations that allocate, free, and reset command buffers or the pool
itself.

Command pools are represented by `VkCommandPool` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCommandPool)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferAllocateInfo.html),
[vkCreateCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCommandPool.html),
[vkDestroyCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCommandPool.html),
[vkFreeCommandBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkFreeCommandBuffers.html),
[vkResetCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetCommandPool.html),
[vkTrimCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html),
[vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPoolKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
