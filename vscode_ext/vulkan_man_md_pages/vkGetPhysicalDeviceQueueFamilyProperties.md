# vkGetPhysicalDeviceQueueFamilyProperties(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyProperties - Reports properties of the
queues of the specified physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query properties of queues available on a physical device, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceQueueFamilyProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pQueueFamilyPropertyCount,
    VkQueueFamilyProperties*                    pQueueFamilyProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the physical device whose properties
  will be queried.

- `pQueueFamilyPropertyCount` is a pointer to an integer related to the
  number of queue families available or queried, as described below.

- `pQueueFamilyProperties` is either `NULL` or a pointer to an array of
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html) structures.

## <a href="#_description" class="anchor"></a>Description

If `pQueueFamilyProperties` is `NULL`, then the number of queue families
available is returned in `pQueueFamilyPropertyCount`. Implementations
**must** support at least one queue family. Otherwise,
`pQueueFamilyPropertyCount` **must** point to a variable set by the user
to the number of elements in the `pQueueFamilyProperties` array, and on
return the variable is overwritten with the number of structures
actually written to `pQueueFamilyProperties`. If
`pQueueFamilyPropertyCount` is less than the number of queue families
available, at most `pQueueFamilyPropertyCount` structures will be
written.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyPropertyCount-parameter  
  `pQueueFamilyPropertyCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyProperties-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyProperties-parameter  
  If the value referenced by `pQueueFamilyPropertyCount` is not `0`, and
  `pQueueFamilyProperties` is not `NULL`, `pQueueFamilyProperties`
  **must** be a valid pointer to an array of `pQueueFamilyPropertyCount`
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
