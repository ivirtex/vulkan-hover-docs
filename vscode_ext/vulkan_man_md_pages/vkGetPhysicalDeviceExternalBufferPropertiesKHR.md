# vkGetPhysicalDeviceExternalBufferProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalBufferProperties - Query external handle
types supported by buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the external handle types supported by buffers, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalBufferProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalBufferInfo*   pExternalBufferInfo,
    VkExternalBufferProperties*                 pExternalBufferProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_external_memory_capabilities
void vkGetPhysicalDeviceExternalBufferPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalBufferInfo*   pExternalBufferInfo,
    VkExternalBufferProperties*                 pExternalBufferProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the buffer
  capabilities.

- `pExternalBufferInfo` is a pointer to a
  [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html)
  structure describing the parameters that would be consumed by
  [vkCreateBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBuffer.html).

- `pExternalBufferProperties` is a pointer to a
  [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)
  structure in which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceExternalBufferProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceExternalBufferProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalBufferProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferInfo-parameter"
  id="VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferInfo-parameter  
  `pExternalBufferInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferProperties-parameter"
  id="VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferProperties-parameter  
  `pExternalBufferProperties` **must** be a valid pointer to a
  [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceExternalBufferProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
