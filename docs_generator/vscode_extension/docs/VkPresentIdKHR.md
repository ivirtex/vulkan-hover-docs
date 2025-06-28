# VkPresentIdKHR(3) Manual Page

## Name

VkPresentIdKHR - The list of presentation identifiers



## [](#_c_specification)C Specification

The `VkPresentIdKHR` structure is defined as:

```c++
// Provided by VK_KHR_present_id
typedef struct VkPresentIdKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           swapchainCount;
    const uint64_t*    pPresentIds;
} VkPresentIdKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is the number of swapchains being presented to the `vkQueuePresentKHR` command.
- `pPresentIds` is `NULL` or a pointer to an array of `uint64_t` with `swapchainCount` entries. If not `NULL`, each non-zero value in `pPresentIds` specifies the present id to be associated with the presentation of the swapchain with the same index in the [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) call.

## [](#_description)Description

For applications to be able to reference specific presentation events queued by a call to `vkQueuePresentKHR`, an identifier needs to be associated with them. When the [`presentId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentId) feature is enabled, applications **can** include the `VkPresentIdKHR` structure in the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure to supply identifiers.

Each `VkSwapchainKHR` has a presentId associated with it. This value is initially zero when the `VkSwapchainKHR` is created.

When a `VkPresentIdKHR` structure with a non-NULL `pPresentIds` is included in the `pNext` chain of a [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure, each `pSwapchains` entry has a presentId associated in the `pPresentIds` array at the same index as the swapchain in the `pSwapchains` array. If this presentId is non-zero, then the application **can** later use this value to refer to that image presentation. A value of zero indicates that this presentation has no associated presentId. A non-zero presentId **must** be greater than any non-zero presentId passed previously by the application for the same swapchain.

There is no requirement for any precise timing relationship between the presentation of the image to the user and the update of the presentId value, but implementations **should** make this as close as possible to the presentation of the first pixel in the new image to the user.

Valid Usage

- [](#VUID-VkPresentIdKHR-swapchainCount-04998)VUID-VkPresentIdKHR-swapchainCount-04998  
  `swapchainCount` **must** be the same value as `VkPresentInfoKHR`::`swapchainCount`, where this `VkPresentIdKHR` is in the `pNext` chain of the `VkPresentInfoKHR` structure
- [](#VUID-VkPresentIdKHR-presentIds-04999)VUID-VkPresentIdKHR-presentIds-04999  
  Each `presentIds` entry **must** be greater than any previous `presentIds` entry passed for the associated `pSwapchains` entry

Valid Usage (Implicit)

- [](#VUID-VkPresentIdKHR-sType-sType)VUID-VkPresentIdKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_ID_KHR`
- [](#VUID-VkPresentIdKHR-pPresentIds-parameter)VUID-VkPresentIdKHR-pPresentIds-parameter  
  If `pPresentIds` is not `NULL`, `pPresentIds` **must** be a valid pointer to an array of `swapchainCount` `uint64_t` values
- [](#VUID-VkPresentIdKHR-swapchainCount-arraylength)VUID-VkPresentIdKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_present\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentIdKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0