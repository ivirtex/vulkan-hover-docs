# vkBindVideoSessionMemoryKHR(3) Manual Page

## Name

vkBindVideoSessionMemoryKHR - Bind Video Memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to a video session object, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkBindVideoSessionMemoryKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    uint32_t                                    bindSessionMemoryInfoCount,
    const VkBindVideoSessionMemoryInfoKHR*      pBindSessionMemoryInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the video session.

- `videoSession` is the video session to be bound with device memory.

- `bindSessionMemoryInfoCount` is the number of elements in
  `pBindSessionMemoryInfos`.

- `pBindSessionMemoryInfos` is a pointer to an array of
  `bindSessionMemoryInfoCount`
  [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindVideoSessionMemoryInfoKHR.html)
  structures specifying memory regions to be bound to specific memory
  bindings of the video session.

## <a href="#_description" class="anchor"></a>Description

The valid usage statements below refer to the
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure
corresponding to a specific element of `pBindSessionMemoryInfos`, which
is defined as follows:

- If the `memoryBindIndex` member of the element of
  `pBindSessionMemoryInfos` in question matches the `memoryBindIndex`
  member of one of the elements returned in `pMemoryRequirements` when
  [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)
  is called with the same `videoSession` and with
  `pMemoryRequirementsCount` equal to `bindSessionMemoryInfoCount`, then
  the `memoryRequirements` member of that element of
  `pMemoryRequirements` is the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure
  corresponding to the element of `pBindSessionMemoryInfos` in question.

- Otherwise the element of `pBindSessionMemoryInfos` in question is said
  to not have a corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure.

Valid Usage

- <a href="#VUID-vkBindVideoSessionMemoryKHR-videoSession-07195"
  id="VUID-vkBindVideoSessionMemoryKHR-videoSession-07195"></a>
  VUID-vkBindVideoSessionMemoryKHR-videoSession-07195  
  The memory binding of `videoSession` identified by the
  `memoryBindIndex` member of any element of `pBindSessionMemoryInfos`
  **must** not already be backed by a memory object

- <a href="#VUID-vkBindVideoSessionMemoryKHR-memoryBindIndex-07196"
  id="VUID-vkBindVideoSessionMemoryKHR-memoryBindIndex-07196"></a>
  VUID-vkBindVideoSessionMemoryKHR-memoryBindIndex-07196  
  The `memoryBindIndex` member of each element of
  `pBindSessionMemoryInfos` **must** be unique within
  `pBindSessionMemoryInfos`

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07197"
  id="VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07197"></a>
  VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07197  
  Each element of `pBindSessionMemoryInfos` **must** have a
  corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)
  structure

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07198"
  id="VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07198"></a>
  VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07198  
  If an element of `pBindSessionMemoryInfos` has a corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure, then the
  `memory` member of that element of `pBindSessionMemoryInfos` **must**
  have been allocated using one of the memory types allowed in the
  `memoryTypeBits` member of the corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07199"
  id="VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07199"></a>
  VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07199  
  If an element of `pBindSessionMemoryInfos` has a corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure, then the
  `memoryOffset` member of that element of `pBindSessionMemoryInfos`
  **must** be an integer multiple of the `alignment` member of the
  corresponding [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)
  structure

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07200"
  id="VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07200"></a>
  VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-07200  
  If an element of `pBindSessionMemoryInfos` has a corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure, then the
  `memorySize` member of that element of `pBindSessionMemoryInfos`
  **must** equal the `size` member of the corresponding
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure

Valid Usage (Implicit)

- <a href="#VUID-vkBindVideoSessionMemoryKHR-device-parameter"
  id="VUID-vkBindVideoSessionMemoryKHR-device-parameter"></a>
  VUID-vkBindVideoSessionMemoryKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindVideoSessionMemoryKHR-videoSession-parameter"
  id="VUID-vkBindVideoSessionMemoryKHR-videoSession-parameter"></a>
  VUID-vkBindVideoSessionMemoryKHR-videoSession-parameter  
  `videoSession` **must** be a valid
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-parameter"
  id="VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-parameter"></a>
  VUID-vkBindVideoSessionMemoryKHR-pBindSessionMemoryInfos-parameter  
  `pBindSessionMemoryInfos` **must** be a valid pointer to an array of
  `bindSessionMemoryInfoCount` valid
  [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindVideoSessionMemoryInfoKHR.html)
  structures

- <a
  href="#VUID-vkBindVideoSessionMemoryKHR-bindSessionMemoryInfoCount-arraylength"
  id="VUID-vkBindVideoSessionMemoryKHR-bindSessionMemoryInfoCount-arraylength"></a>
  VUID-vkBindVideoSessionMemoryKHR-bindSessionMemoryInfoCount-arraylength  
  `bindSessionMemoryInfoCount` **must** be greater than `0`

- <a href="#VUID-vkBindVideoSessionMemoryKHR-videoSession-parent"
  id="VUID-vkBindVideoSessionMemoryKHR-videoSession-parent"></a>
  VUID-vkBindVideoSessionMemoryKHR-videoSession-parent  
  `videoSession` **must** have been created, allocated, or retrieved
  from `device`

Host Synchronization

- Host access to `videoSession` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindVideoSessionMemoryInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindVideoSessionMemoryKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
