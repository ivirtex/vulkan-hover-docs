# VkVideoSessionKHR(3) Manual Page

## Name

VkVideoSessionKHR - Opaque handle to a video session object



## [](#_c_specification)C Specification

Video sessions are objects that represent and maintain the state needed to perform video decode or encode operations using a specific video profile.

In case of video encode profiles this includes the current [rate control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) configuration and the currently set [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level).

Video sessions are represented by `VkVideoSessionKHR` handles:

```c++
// Provided by VK_KHR_video_queue
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkVideoSessionKHR)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html), [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html), [vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindVideoSessionMemoryKHR.html), [vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateVideoSessionKHR.html), [vkDestroyVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyVideoSessionKHR.html), [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0