# VK\_LUID\_SIZE(3) Manual Page

## Name

VK\_LUID\_SIZE - Length of a locally unique device identifier



## [](#_c_specification)C Specification

`VK_LUID_SIZE` is the length in `uint8_t` values of an array containing a locally unique device identifier, as returned in [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html)::`deviceLUID`.

```c++
#define VK_LUID_SIZE                      8U
```

or the equivalent

```c++
#define VK_LUID_SIZE_KHR                  VK_LUID_SIZE
```

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_capabilities.html), [VK\_KHR\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_capabilities.html), [VK\_KHR\_external\_semaphore\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_capabilities.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_LUID_SIZE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0