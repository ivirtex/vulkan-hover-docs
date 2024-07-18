# VkPresentIdKHR(3) Manual Page

## Name

VkPresentIdKHR - The list of presentation identifiers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPresentIdKHR` structure is defined as:

``` c
// Provided by VK_KHR_present_id
typedef struct VkPresentIdKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           swapchainCount;
    const uint64_t*    pPresentIds;
} VkPresentIdKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchainCount` is the number of swapchains being presented to the
  `vkQueuePresentKHR` command.

- `pPresentIds` is `NULL` or a pointer to an array of `uint64_t` with
  `swapchainCount` entries. If not `NULL`, each non-zero value in
  `pPresentIds` specifies the present id to be associated with the
  presentation of the swapchain with the same index in the
  [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) call.

## <a href="#_description" class="anchor"></a>Description

For applications to be able to reference specific presentation events
queued by a call to `vkQueuePresentKHR`, an identifier needs to be
associated with them. When the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-presentId"
target="_blank" rel="noopener"><code>presentId</code></a> feature is
enabled, applications **can** include the `VkPresentIdKHR` structure in
the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html)
structure to supply identifiers.

Each `VkSwapchainKHR` has a presentId associated with it. This value is
initially set to zero when the `VkSwapchainKHR` is created.

When a `VkPresentIdKHR` structure with a non-NULL `pPresentIds` is
included in the `pNext` chain of a
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) structure, each `pSwapchains`
entry has a presentId associated in the `pPresentIds` array at the same
index as the swapchain in the `pSwapchains` array. If this presentId is
non-zero, then the application **can** later use this value to refer to
that image presentation. A value of zero indicates that this
presentation has no associated presentId. A non-zero presentId **must**
be greater than any non-zero presentId passed previously by the
application for the same swapchain.

There is no requirement for any precise timing relationship between the
presentation of the image to the user and the update of the presentId
value, but implementations **should** make this as close as possible to
the presentation of the first pixel in the new image to the user.

Valid Usage

- <a href="#VUID-VkPresentIdKHR-swapchainCount-04998"
  id="VUID-VkPresentIdKHR-swapchainCount-04998"></a>
  VUID-VkPresentIdKHR-swapchainCount-04998  
  `swapchainCount` **must** be the same value as
  `VkPresentInfoKHR`::`swapchainCount`, where this `VkPresentIdKHR` is
  in the `pNext` chain of the `VkPresentInfoKHR` structure

- <a href="#VUID-VkPresentIdKHR-presentIds-04999"
  id="VUID-VkPresentIdKHR-presentIds-04999"></a>
  VUID-VkPresentIdKHR-presentIds-04999  
  Each `presentIds` entry **must** be greater than any previous
  `presentIds` entry passed for the associated `pSwapchains` entry

Valid Usage (Implicit)

- <a href="#VUID-VkPresentIdKHR-sType-sType"
  id="VUID-VkPresentIdKHR-sType-sType"></a>
  VUID-VkPresentIdKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_ID_KHR`

- <a href="#VUID-VkPresentIdKHR-pPresentIds-parameter"
  id="VUID-VkPresentIdKHR-pPresentIds-parameter"></a>
  VUID-VkPresentIdKHR-pPresentIds-parameter  
  If `pPresentIds` is not `NULL`, `pPresentIds` **must** be a valid
  pointer to an array of `swapchainCount` `uint64_t` values

- <a href="#VUID-VkPresentIdKHR-swapchainCount-arraylength"
  id="VUID-VkPresentIdKHR-swapchainCount-arraylength"></a>
  VUID-VkPresentIdKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_present_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_present_id.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentIdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
