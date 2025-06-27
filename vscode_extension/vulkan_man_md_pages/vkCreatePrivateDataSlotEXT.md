# vkCreatePrivateDataSlot(3) Manual Page

## Name

vkCreatePrivateDataSlot - Create a slot for private data storage



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a private data slot, call:

``` c
// Provided by VK_VERSION_1_3
VkResult vkCreatePrivateDataSlot(
    VkDevice                                    device,
    const VkPrivateDataSlotCreateInfo*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPrivateDataSlot*                          pPrivateDataSlot);
```

or the equivalent command

``` c
// Provided by VK_EXT_private_data
VkResult vkCreatePrivateDataSlotEXT(
    VkDevice                                    device,
    const VkPrivateDataSlotCreateInfo*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPrivateDataSlot*                          pPrivateDataSlot);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device associated with the creation of the
  object(s) holding the private data slot.

- `pCreateInfo` is a pointer to a `VkPrivateDataSlotCreateInfo`

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPrivateDataSlot` is a pointer to a
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) handle in which the
  resulting private data slot is returned

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreatePrivateDataSlot-privateData-04564"
  id="VUID-vkCreatePrivateDataSlot-privateData-04564"></a>
  VUID-vkCreatePrivateDataSlot-privateData-04564  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-privateData"
  target="_blank" rel="noopener"><code>privateData</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCreatePrivateDataSlot-device-parameter"
  id="VUID-vkCreatePrivateDataSlot-device-parameter"></a>
  VUID-vkCreatePrivateDataSlot-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreatePrivateDataSlot-pCreateInfo-parameter"
  id="VUID-vkCreatePrivateDataSlot-pCreateInfo-parameter"></a>
  VUID-vkCreatePrivateDataSlot-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateInfo.html)
  structure

- <a href="#VUID-vkCreatePrivateDataSlot-pAllocator-parameter"
  id="VUID-vkCreatePrivateDataSlot-pAllocator-parameter"></a>
  VUID-vkCreatePrivateDataSlot-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreatePrivateDataSlot-pPrivateDataSlot-parameter"
  id="VUID-vkCreatePrivateDataSlot-pPrivateDataSlot-parameter"></a>
  VUID-vkCreatePrivateDataSlot-pPrivateDataSlot-parameter  
  `pPrivateDataSlot` **must** be a valid pointer to a
  [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html),
[VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreatePrivateDataSlot"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
