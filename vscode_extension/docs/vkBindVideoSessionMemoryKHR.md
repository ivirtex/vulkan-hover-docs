# vkBindVideoSessionMemoryKHR(3) Manual Page

## Name

vkBindVideoSessionMemoryKHR - Bind Video Memory



## [](#_c_specification)C Specification

To attach memory to a video session object, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkBindVideoSessionMemoryKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    uint32_t                                    bindSessionMemoryInfoCount,
    const VkBindVideoSessionMemoryInfoKHR*      pBindSessionMemoryInfos);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the video session.
- `videoSession` is the video session to be bound with device memory.
- `bindSessionMemoryInfoCount` is the number of elements in `pBindSessionMemoryInfos`.
- `pBindSessionMemoryInfos` is a pointer to an array of `bindSessionMemoryInfoCount` [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVideoSessionMemoryInfoKHR.html) structures specifying memory regions to be bound to specific memory bindings of the video session.

## [](#_description)Description

The valid usage statements below refer to the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure corresponding to a specific element of `pBindSessionMemoryInfos`, which is defined as follows:

- If the `memoryBindIndex` member of the element of `pBindSessionMemoryInfos` in question matches the `memoryBindIndex` member of one of the elements returned in `pMemoryRequirements` when [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetVideoSessionMemoryRequirementsKHR.html) is called with the same `videoSession` and with `pMemoryRequirementsCount` equal to `bindSessionMemoryInfoCount`, then the `memoryRequirements` member of that element of `pMemoryRequirements` is the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure corresponding to the element of `pBindSessionMemoryInfos` in question.
- Otherwise the element of `pBindSessionMemoryInfos` in question is said to not have a corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure.

Valid Usage

- [](#VUID-vkBindVideoSessionMemoryKHR-videoSession-07195)VUID-vkBindVideoSessionMemoryKHR-videoSession-07195  
  The memory binding of `videoSession` identified by the `memoryBindIndex` member of any element of `pBindSessionMemoryInfos` **must** not already be backed by a memory object
- [](#VUID-vkBindVideoSessionMemoryKHR-memoryBindIndex-07196)VUID-vkBindVideoSessionMemoryKHR-memoryBindIndex-07196  
  The `memoryBindIndex` member of each element of `pBindSessionMemoryInfos` **must** be unique within `pBindSessionMemoryInfos`
- [](#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07197)VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07197  
  Each element of `pBindSessionMemoryInfos` **must** have a corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure
- [](#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07198)VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07198  
  If an element of `pBindSessionMemoryInfos` has a corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure, then the `memory` member of that element of `pBindSessionMemoryInfos` **must** have been allocated using one of the memory types allowed in the `memoryTypeBits` member of the corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure
- [](#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07199)VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07199  
  If an element of `pBindSessionMemoryInfos` has a corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure, then the `memoryOffset` member of that element of `pBindSessionMemoryInfos` **must** be an integer multiple of the `alignment` member of the corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure
- [](#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07200)VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07200  
  If an element of `pBindSessionMemoryInfos` has a corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure, then the `memorySize` member of that element of `pBindSessionMemoryInfos` **must** equal the `size` member of the corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure

Valid Usage (Implicit)

- [](#VUID-vkBindVideoSessionMemoryKHR-device-parameter)VUID-vkBindVideoSessionMemoryKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkBindVideoSessionMemoryKHR-videoSession-parameter)VUID-vkBindVideoSessionMemoryKHR-videoSession-parameter  
  `videoSession` **must** be a valid [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html) handle
- [](#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-parameter)VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-parameter  
  `pBindSessionMemoryInfos` **must** be a valid pointer to an array of `bindSessionMemoryInfoCount` valid [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVideoSessionMemoryInfoKHR.html) structures
- [](#VUID-vkBindVideoSessionMemoryKHR-bindSessionMemoryInfoCount-arraylength)VUID-vkBindVideoSessionMemoryKHR-bindSessionMemoryInfoCount-arraylength  
  `bindSessionMemoryInfoCount` **must** be greater than `0`
- [](#VUID-vkBindVideoSessionMemoryKHR-videoSession-parent)VUID-vkBindVideoSessionMemoryKHR-videoSession-parent  
  `videoSession` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `videoSession` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVideoSessionMemoryInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkBindVideoSessionMemoryKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0