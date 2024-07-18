# vkGetBufferMemoryRequirements2(3) Manual Page

## Name

vkGetBufferMemoryRequirements2 - Returns the memory requirements for
specified Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for a buffer resource, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetBufferMemoryRequirements2(
    VkDevice                                    device,
    const VkBufferMemoryRequirementsInfo2*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_memory_requirements2
void vkGetBufferMemoryRequirements2KHR(
    VkDevice                                    device,
    const VkBufferMemoryRequirementsInfo2*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffer.

- `pInfo` is a pointer to a
  [VkBufferMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html)
  structure containing parameters required for the memory requirements
  query.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure in which
  the memory requirements of the buffer object are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetBufferMemoryRequirements2-device-parameter"
  id="VUID-vkGetBufferMemoryRequirements2-device-parameter"></a>
  VUID-vkGetBufferMemoryRequirements2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetBufferMemoryRequirements2-pInfo-parameter"
  id="VUID-vkGetBufferMemoryRequirements2-pInfo-parameter"></a>
  VUID-vkGetBufferMemoryRequirements2-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkBufferMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html)
  structure

- <a
  href="#VUID-vkGetBufferMemoryRequirements2-pMemoryRequirements-parameter"
  id="VUID-vkGetBufferMemoryRequirements2-pMemoryRequirements-parameter"></a>
  VUID-vkGetBufferMemoryRequirements2-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkBufferMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferMemoryRequirements2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
