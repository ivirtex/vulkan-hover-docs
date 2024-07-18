# vkGetDeviceImageSparseMemoryRequirements(3) Manual Page

## Name

vkGetDeviceImageSparseMemoryRequirements - Query the memory requirements
for a sparse image



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the sparse memory requirements for an image resource
without creating an object, call:

``` c
// Provided by VK_VERSION_1_3
void vkGetDeviceImageSparseMemoryRequirements(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

or the equivalent command

``` c
// Provided by VK_KHR_maintenance4
void vkGetDeviceImageSparseMemoryRequirementsKHR(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device intended to own the image.

- `pInfo` is a pointer to a
  [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html)
  structure containing parameters required for the memory requirements
  query.

- `pSparseMemoryRequirementCount` is a pointer to an integer related to
  the number of sparse memory requirements available or queried, as
  described below.

- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array
  of `VkSparseImageMemoryRequirements2` structures.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDeviceImageSparseMemoryRequirements-device-parameter"
  id="VUID-vkGetDeviceImageSparseMemoryRequirements-device-parameter"></a>
  VUID-vkGetDeviceImageSparseMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceImageSparseMemoryRequirements-pInfo-parameter"
  id="VUID-vkGetDeviceImageSparseMemoryRequirements-pInfo-parameter"></a>
  VUID-vkGetDeviceImageSparseMemoryRequirements-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html)
  structure

- <a
  href="#VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter"
  id="VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter"></a>
  VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter"
  id="VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter"></a>
  VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`,
  and `pSparseMemoryRequirements` is not `NULL`,
  `pSparseMemoryRequirements` **must** be a valid pointer to an array of
  `pSparseMemoryRequirementCount`
  [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html),
[VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceImageSparseMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
