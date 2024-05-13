# vkGetPhysicalDeviceExternalFenceProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalFenceProperties - Function for querying
external fence handle capabilities.



## <a href="#_c_specification" class="anchor"></a>C Specification

Fences **may** support import and export of their <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-payloads"
target="_blank" rel="noopener">payload</a> to external handles. To query
the external handle types supported by fences, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalFenceProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalFenceInfo*    pExternalFenceInfo,
    VkExternalFenceProperties*                  pExternalFenceProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_external_fence_capabilities
void vkGetPhysicalDeviceExternalFencePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalFenceInfo*    pExternalFenceInfo,
    VkExternalFenceProperties*                  pExternalFenceProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the fence
  capabilities.

- `pExternalFenceInfo` is a pointer to a
  [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)
  structure describing the parameters that would be consumed by
  [vkCreateFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html).

- `pExternalFenceProperties` is a pointer to a
  [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html) structure
  in which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceExternalFenceProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceExternalFenceProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalFenceProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceInfo-parameter"
  id="VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceInfo-parameter  
  `pExternalFenceInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceProperties-parameter"
  id="VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceProperties-parameter  
  `pExternalFenceProperties` **must** be a valid pointer to a
  [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceExternalFenceProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
