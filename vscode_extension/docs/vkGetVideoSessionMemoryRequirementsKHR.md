# vkGetVideoSessionMemoryRequirementsKHR(3) Manual Page

## Name

vkGetVideoSessionMemoryRequirementsKHR - Get the memory requirements for
a video session



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for a video session object, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkGetVideoSessionMemoryRequirementsKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    uint32_t*                                   pMemoryRequirementsCount,
    VkVideoSessionMemoryRequirementsKHR*        pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the video session.

- `videoSession` is the video session to query.

- `pMemoryRequirementsCount` is a pointer to an integer related to the
  number of memory binding requirements available or queried, as
  described below.

- `pMemoryRequirements` is `NULL` or a pointer to an array of
  [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionMemoryRequirementsKHR.html)
  structures in which the memory binding requirements of the video
  session are returned.

## <a href="#_description" class="anchor"></a>Description

If `pMemoryRequirements` is `NULL`, then the number of memory bindings
required for the video session is returned in
`pMemoryRequirementsCount`. Otherwise, `pMemoryRequirementsCount`
**must** point to a variable set by the application to the number of
elements in the `pMemoryRequirements` array, and on return the variable
is overwritten with the number of memory binding requirements actually
written to `pMemoryRequirements`. If `pMemoryRequirementsCount` is less
than the number of memory bindings required for the video session, then
at most `pMemoryRequirementsCount` elements will be written to
`pMemoryRequirements`, and `VK_INCOMPLETE` will be returned, instead of
`VK_SUCCESS`, to indicate that not all required memory binding
requirements were returned.

Valid Usage (Implicit)

- <a href="#VUID-vkGetVideoSessionMemoryRequirementsKHR-device-parameter"
  id="VUID-vkGetVideoSessionMemoryRequirementsKHR-device-parameter"></a>
  VUID-vkGetVideoSessionMemoryRequirementsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parameter"
  id="VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parameter"></a>
  VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parameter  
  `videoSession` **must** be a valid
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

- <a
  href="#VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirementsCount-parameter"
  id="VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirementsCount-parameter"></a>
  VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirementsCount-parameter  
  `pMemoryRequirementsCount` **must** be a valid pointer to a `uint32_t`
  value

- <a
  href="#VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirements-parameter"
  id="VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirements-parameter"></a>
  VUID-vkGetVideoSessionMemoryRequirementsKHR-pMemoryRequirements-parameter  
  If the value referenced by `pMemoryRequirementsCount` is not `0`, and
  `pMemoryRequirements` is not `NULL`, `pMemoryRequirements` **must** be
  a valid pointer to an array of `pMemoryRequirementsCount`
  [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionMemoryRequirementsKHR.html)
  structures

- <a
  href="#VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parent"
  id="VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parent"></a>
  VUID-vkGetVideoSessionMemoryRequirementsKHR-videoSession-parent  
  `videoSession` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

- `VK_INCOMPLETE`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html),
[VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionMemoryRequirementsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetVideoSessionMemoryRequirementsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
