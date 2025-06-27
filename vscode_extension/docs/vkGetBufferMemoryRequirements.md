# vkGetBufferMemoryRequirements(3) Manual Page

## Name

vkGetBufferMemoryRequirements - Returns the memory requirements for
specified Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for a buffer resource, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetBufferMemoryRequirements(
    VkDevice                                    device,
    VkBuffer                                    buffer,
    VkMemoryRequirements*                       pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffer.

- `buffer` is the buffer to query.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure in which
  the memory requirements of the buffer object are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetBufferMemoryRequirements-device-parameter"
  id="VUID-vkGetBufferMemoryRequirements-device-parameter"></a>
  VUID-vkGetBufferMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetBufferMemoryRequirements-buffer-parameter"
  id="VUID-vkGetBufferMemoryRequirements-buffer-parameter"></a>
  VUID-vkGetBufferMemoryRequirements-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a
  href="#VUID-vkGetBufferMemoryRequirements-pMemoryRequirements-parameter"
  id="VUID-vkGetBufferMemoryRequirements-pMemoryRequirements-parameter"></a>
  VUID-vkGetBufferMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure

- <a href="#VUID-vkGetBufferMemoryRequirements-buffer-parent"
  id="VUID-vkGetBufferMemoryRequirements-buffer-parent"></a>
  VUID-vkGetBufferMemoryRequirements-buffer-parent  
  `buffer` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
