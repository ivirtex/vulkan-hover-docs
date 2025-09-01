# vkGetVideoSessionMemoryRequirementsKHR(3) Manual Page

## Name

vkGetVideoSessionMemoryRequirementsKHR - Get the memory requirements for a video session



## [](#_c_specification)C Specification

To determine the memory requirements for a video session object, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkGetVideoSessionMemoryRequirementsKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    uint32_t*                                   pMemoryRequirementsCount,
    VkVideoSessionMemoryRequirementsKHR*        pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the video session.
- `videoSession` is the video session to query.
- `pMemoryRequirementsCount` is a pointer to an integer related to the number of memory binding requirements available or queried, as described below.
- `pMemoryRequirements` is `NULL` or a pointer to an array of [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionMemoryRequirementsKHR.html) structures in which the memory binding requirements of the video session are returned.

## [](#_description)Description

If `pMemoryRequirements` is `NULL`, then the number of memory bindings required for the video session is returned in `pMemoryRequirementsCount`. Otherwise, `pMemoryRequirementsCount` **must** point to a variable set by the application to the number of elements in the `pMemoryRequirements` array, and on return the variable is overwritten with the number of memory binding requirements actually written to `pMemoryRequirements`. If `pMemoryRequirementsCount` is less than the number of memory bindings required for the video session, then at most `pMemoryRequirementsCount` elements will be written to `pMemoryRequirements`, and `VK_INCOMPLETE` will be returned, instead of `VK_SUCCESS`, to indicate that not all required memory binding requirements were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetVideoSessionMemoryRequirementsKHR-device-parameter)VUID-vkGetVideoSessionMemoryRequirementsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parameter)VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parameter  
  `videoSession` **must** be a valid [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html) handle
- [](#VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirementsCount-parameter)VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirementsCount-parameter  
  `pMemoryRequirementsCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirements-parameter)VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirements-parameter  
  If the value referenced by `pMemoryRequirementsCount` is not `0`, and `pMemoryRequirements` is not `NULL`, `pMemoryRequirements` **must** be a valid pointer to an array of `pMemoryRequirementsCount` [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionMemoryRequirementsKHR.html) structures
- [](#VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parent)VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parent  
  `videoSession` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html), [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionMemoryRequirementsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetVideoSessionMemoryRequirementsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0