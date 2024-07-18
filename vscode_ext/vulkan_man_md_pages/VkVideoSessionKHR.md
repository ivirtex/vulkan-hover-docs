# VkVideoSessionKHR(3) Manual Page

## Name

VkVideoSessionKHR - Opaque handle to a video session object



## <a href="#_c_specification" class="anchor"></a>C Specification

Video sessions are objects that represent and maintain the state needed
to perform video decode or encode operations using a specific video
profile.

In case of video encode profiles this includes the current <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
target="_blank" rel="noopener">rate control</a> configuration and the
currently set <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
target="_blank" rel="noopener">video encode quality level</a>.

Video sessions are represented by `VkVideoSessionKHR` handles:

``` c
// Provided by VK_KHR_video_queue
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkVideoSessionKHR)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html),
[VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html),
[vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindVideoSessionMemoryKHR.html),
[vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionKHR.html),
[vkDestroyVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyVideoSessionKHR.html),
[vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
