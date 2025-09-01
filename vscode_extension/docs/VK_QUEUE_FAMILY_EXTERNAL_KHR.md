# VK\_QUEUE\_FAMILY\_EXTERNAL(3) Manual Page

## Name

VK\_QUEUE\_FAMILY\_EXTERNAL - External queue family index sentinel



## [](#_c_specification)C Specification

The special queue family index `VK_QUEUE_FAMILY_EXTERNAL` represents any queue external to the resource’s current Vulkan instance, as long as the queue uses the same underlying device group or physical device, and the same driver version as the resource’s [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), as indicated by [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html)::`deviceUUID` and [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html)::`driverUUID`.

```c++
#define VK_QUEUE_FAMILY_EXTERNAL          (~1U)
```

or the equivalent

```c++
#define VK_QUEUE_FAMILY_EXTERNAL_KHR      VK_QUEUE_FAMILY_EXTERNAL
```

## [](#_see_also)See Also

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QUEUE_FAMILY_EXTERNAL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0