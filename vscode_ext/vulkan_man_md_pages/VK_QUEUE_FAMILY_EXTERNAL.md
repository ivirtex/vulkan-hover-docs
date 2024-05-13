# VK_QUEUE_FAMILY_EXTERNAL(3) Manual Page

## Name

VK_QUEUE_FAMILY_EXTERNAL - External queue family index sentinel



## <a href="#_c_specification" class="anchor"></a>C Specification

The special queue family index `VK_QUEUE_FAMILY_EXTERNAL` represents any
queue external to the resource’s current Vulkan instance, as long as the
queue uses the same underlying device group or physical device, and the
same driver version as the resource’s [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), as
indicated by
[VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html)::`deviceUUID`
and
[VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html)::`driverUUID`.

``` c
#define VK_QUEUE_FAMILY_EXTERNAL          (~1U)
```

or the equivalent

``` c
#define VK_QUEUE_FAMILY_EXTERNAL_KHR      VK_QUEUE_FAMILY_EXTERNAL
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QUEUE_FAMILY_EXTERNAL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
