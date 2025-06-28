# vkCreatePrivateDataSlot(3) Manual Page

## Name

vkCreatePrivateDataSlot - Create a slot for private data storage



## [](#_c_specification)C Specification

To create a private data slot, call:

```c++
// Provided by VK_VERSION_1_3
VkResult vkCreatePrivateDataSlot(
    VkDevice                                    device,
    const VkPrivateDataSlotCreateInfo*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPrivateDataSlot*                          pPrivateDataSlot);
```

or the equivalent command

```c++
// Provided by VK_EXT_private_data
VkResult vkCreatePrivateDataSlotEXT(
    VkDevice                                    device,
    const VkPrivateDataSlotCreateInfo*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPrivateDataSlot*                          pPrivateDataSlot);
```

## [](#_parameters)Parameters

- `device` is the logical device associated with the creation of the object(s) holding the private data slot.
- `pCreateInfo` is a pointer to a [VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateInfo.html)
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPrivateDataSlot` is a pointer to a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) handle in which the resulting private data slot is returned

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreatePrivateDataSlot-privateData-04564)VUID-vkCreatePrivateDataSlot-privateData-04564  
  The [`privateData`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-privateData) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreatePrivateDataSlot-device-parameter)VUID-vkCreatePrivateDataSlot-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreatePrivateDataSlot-pCreateInfo-parameter)VUID-vkCreatePrivateDataSlot-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateInfo.html) structure
- [](#VUID-vkCreatePrivateDataSlot-pAllocator-parameter)VUID-vkCreatePrivateDataSlot-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreatePrivateDataSlot-pPrivateDataSlot-parameter)VUID-vkCreatePrivateDataSlot-pPrivateDataSlot-parameter  
  `pPrivateDataSlot` **must** be a valid pointer to a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) handle
- [](#VUID-vkCreatePrivateDataSlot-device-queuecount)VUID-vkCreatePrivateDataSlot-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html), [VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreatePrivateDataSlot)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0