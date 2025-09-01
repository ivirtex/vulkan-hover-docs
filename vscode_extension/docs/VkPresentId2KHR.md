# VkPresentId2KHR(3) Manual Page

## Name

VkPresentId2KHR - The list of presentation identifiers



## [](#_c_specification)C Specification

The `VkPresentId2KHR` structure is defined as:

```c++
// Provided by VK_KHR_present_id2
typedef struct VkPresentId2KHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           swapchainCount;
    const uint64_t*    pPresentIds;
} VkPresentId2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is the number of swapchains being presented to the `vkQueuePresentKHR` command.
- `pPresentIds` is `NULL` or a pointer to an array of uint64\_t with `swapchainCount` entries. If not `NULL`, each non-zero value in `pPresentIds` specifies the present id to be associated with the presentation of the swapchain with the same index in the [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) call.

## [](#_description)Description

For applications to be able to reference specific presentation events queued by a call to `vkQueuePresentKHR`, an identifier needs to be associated with them.

When the [VkSurfaceCapabilitiesPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentId2KHR.html) surface capability is present for a surface, applications **can** include the `VkPresentId2KHR` structure in the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure to associate an identifier with each presentation request. The `pPresentIds` provides an identifier for the swapchain present at the corresponding index in [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html)’s `pSwapchains` array.

If this presentId is non-zero, then the application **can** later use this value to refer to that image presentation. A value of zero indicates that this presentation has no associated presentId. A non-zero presentId **must** be greater than any non-zero presentId passed previously by the application for the same swapchain.

If a non-zero presentId was provided, this may be used with [vkWaitForPresent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresent2KHR.html) for the application to synchronize against the presentation engine’s processing of the presentation request.

Note

The ID namespace used by this extension **must** be shared with other extensions that allow the application to provide a 64-bit monotonically increasing presentation ID, such as the original VK\_KHR\_present\_id.

This is to allow existing extensions that depend on VK\_KHR\_present\_id to use VK\_KHR\_present\_id2 provided IDs without change, as well as to simplify writing future extensions that require application provided presentation IDs.

Valid Usage

- [](#VUID-VkPresentId2KHR-swapchainCount-10818)VUID-VkPresentId2KHR-swapchainCount-10818  
  `swapchainCount` **must** be the same value as `VkPresentInfoKHR`::`swapchainCount`, where this `VkPresentId2KHR` is in the `pNext` chain of the `VkPresentInfoKHR` structure
- [](#VUID-VkPresentId2KHR-presentIds-10819)VUID-VkPresentId2KHR-presentIds-10819  
  Each `presentIds` entry **must** be greater than any previous `presentIds` entry passed for the associated `pSwapchains` entry
- [](#VUID-VkPresentId2KHR-None-10820)VUID-VkPresentId2KHR-None-10820  
  The swapchain must have been created with `VK_SWAPCHAIN_CREATE_PRESENT_ID_2_BIT_KHR` bit set in the `VkSwapchainCreateFlagBitsKHR` field

Valid Usage (Implicit)

- [](#VUID-VkPresentId2KHR-sType-sType)VUID-VkPresentId2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_ID_2_KHR`
- [](#VUID-VkPresentId2KHR-pPresentIds-parameter)VUID-VkPresentId2KHR-pPresentIds-parameter  
  If `pPresentIds` is not `NULL`, `pPresentIds` **must** be a valid pointer to an array of `swapchainCount` `uint64_t` values
- [](#VUID-VkPresentId2KHR-swapchainCount-arraylength)VUID-VkPresentId2KHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_present\_id2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentId2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0