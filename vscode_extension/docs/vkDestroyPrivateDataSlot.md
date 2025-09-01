# vkDestroyPrivateDataSlot(3) Manual Page

## Name

vkDestroyPrivateDataSlot - Destroy a private data slot



## [](#_c_specification)C Specification

To destroy a private data slot, call:

```c++
// Provided by VK_VERSION_1_3
void vkDestroyPrivateDataSlot(
    VkDevice                                    device,
    VkPrivateDataSlot                           privateDataSlot,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

```c++
// Provided by VK_EXT_private_data
void vkDestroyPrivateDataSlotEXT(
    VkDevice                                    device,
    VkPrivateDataSlot                           privateDataSlot,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device associated with the creation of the object(s) holding the private data slot.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `privateDataSlot` is the private data slot to destroy.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyPrivateDataSlot-privateDataSlot-04062)VUID-vkDestroyPrivateDataSlot-privateDataSlot-04062  
  If `VkAllocationCallbacks` were provided when `privateDataSlot` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyPrivateDataSlot-privateDataSlot-04063)VUID-vkDestroyPrivateDataSlot-privateDataSlot-04063  
  If no `VkAllocationCallbacks` were provided when `privateDataSlot` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyPrivateDataSlot-device-parameter)VUID-vkDestroyPrivateDataSlot-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyPrivateDataSlot-privateDataSlot-parameter)VUID-vkDestroyPrivateDataSlot-privateDataSlot-parameter  
  If `privateDataSlot` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `privateDataSlot` **must** be a valid [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) handle
- [](#VUID-vkDestroyPrivateDataSlot-pAllocator-parameter)VUID-vkDestroyPrivateDataSlot-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyPrivateDataSlot-privateDataSlot-parent)VUID-vkDestroyPrivateDataSlot-privateDataSlot-parent  
  If `privateDataSlot` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `privateDataSlot` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyPrivateDataSlot).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0