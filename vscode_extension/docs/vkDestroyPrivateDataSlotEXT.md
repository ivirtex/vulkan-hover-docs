# vkDestroyPrivateDataSlot(3) Manual Page

## Name

vkDestroyPrivateDataSlot - Destroy a private data slot



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a private data slot, call:

``` c
// Provided by VK_VERSION_1_3
void vkDestroyPrivateDataSlot(
    VkDevice                                    device,
    VkPrivateDataSlot                           privateDataSlot,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

``` c
// Provided by VK_EXT_private_data
void vkDestroyPrivateDataSlotEXT(
    VkDevice                                    device,
    VkPrivateDataSlot                           privateDataSlot,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device associated with the creation of the
  object(s) holding the private data slot.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `privateDataSlot` is the private data slot to destroy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyPrivateDataSlot-privateDataSlot-04062"
  id="VUID-vkDestroyPrivateDataSlot-privateDataSlot-04062"></a>
  VUID-vkDestroyPrivateDataSlot-privateDataSlot-04062  
  If `VkAllocationCallbacks` were provided when `privateDataSlot` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyPrivateDataSlot-privateDataSlot-04063"
  id="VUID-vkDestroyPrivateDataSlot-privateDataSlot-04063"></a>
  VUID-vkDestroyPrivateDataSlot-privateDataSlot-04063  
  If no `VkAllocationCallbacks` were provided when `privateDataSlot` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyPrivateDataSlot-device-parameter"
  id="VUID-vkDestroyPrivateDataSlot-device-parameter"></a>
  VUID-vkDestroyPrivateDataSlot-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyPrivateDataSlot-privateDataSlot-parameter"
  id="VUID-vkDestroyPrivateDataSlot-privateDataSlot-parameter"></a>
  VUID-vkDestroyPrivateDataSlot-privateDataSlot-parameter  
  If `privateDataSlot` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `privateDataSlot` **must** be a valid
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) handle

- <a href="#VUID-vkDestroyPrivateDataSlot-pAllocator-parameter"
  id="VUID-vkDestroyPrivateDataSlot-pAllocator-parameter"></a>
  VUID-vkDestroyPrivateDataSlot-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyPrivateDataSlot-privateDataSlot-parent"
  id="VUID-vkDestroyPrivateDataSlot-privateDataSlot-parent"></a>
  VUID-vkDestroyPrivateDataSlot-privateDataSlot-parent  
  If `privateDataSlot` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `privateDataSlot` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyPrivateDataSlot"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
