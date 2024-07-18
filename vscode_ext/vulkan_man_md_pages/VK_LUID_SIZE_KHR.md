# VK_LUID_SIZE(3) Manual Page

## Name

VK_LUID_SIZE - Length of a locally unique device identifier



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_LUID_SIZE` is the length in `uint8_t` values of an array containing
a locally unique device identifier, as returned in
[VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html)::`deviceLUID`.

``` c
#define VK_LUID_SIZE                      8U
```

or the equivalent

``` c
#define VK_LUID_SIZE_KHR                  VK_LUID_SIZE
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_fence_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_capabilities.html),
[VK_KHR_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html),
[VK_KHR_external_semaphore_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_capabilities.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_LUID_SIZE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
