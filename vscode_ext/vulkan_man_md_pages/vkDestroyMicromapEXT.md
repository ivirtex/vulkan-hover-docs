# vkDestroyMicromapEXT(3) Manual Page

## Name

vkDestroyMicromapEXT - Destroy a micromap object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a micromap, call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkDestroyMicromapEXT(
    VkDevice                                    device,
    VkMicromapEXT                               micromap,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the micromap.

- `micromap` is the micromap to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyMicromapEXT-micromap-07441"
  id="VUID-vkDestroyMicromapEXT-micromap-07441"></a>
  VUID-vkDestroyMicromapEXT-micromap-07441  
  All submitted commands that refer to `micromap` **must** have
  completed execution

- <a href="#VUID-vkDestroyMicromapEXT-micromap-07442"
  id="VUID-vkDestroyMicromapEXT-micromap-07442"></a>
  VUID-vkDestroyMicromapEXT-micromap-07442  
  If `VkAllocationCallbacks` were provided when `micromap` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyMicromapEXT-micromap-07443"
  id="VUID-vkDestroyMicromapEXT-micromap-07443"></a>
  VUID-vkDestroyMicromapEXT-micromap-07443  
  If no `VkAllocationCallbacks` were provided when `micromap` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyMicromapEXT-device-parameter"
  id="VUID-vkDestroyMicromapEXT-device-parameter"></a>
  VUID-vkDestroyMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyMicromapEXT-micromap-parameter"
  id="VUID-vkDestroyMicromapEXT-micromap-parameter"></a>
  VUID-vkDestroyMicromapEXT-micromap-parameter  
  If `micromap` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `micromap`
  **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-vkDestroyMicromapEXT-pAllocator-parameter"
  id="VUID-vkDestroyMicromapEXT-pAllocator-parameter"></a>
  VUID-vkDestroyMicromapEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyMicromapEXT-micromap-parent"
  id="VUID-vkDestroyMicromapEXT-micromap-parent"></a>
  VUID-vkDestroyMicromapEXT-micromap-parent  
  If `micromap` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `micromap` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
