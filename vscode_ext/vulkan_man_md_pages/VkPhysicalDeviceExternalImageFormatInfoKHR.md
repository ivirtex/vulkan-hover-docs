# VkPhysicalDeviceExternalImageFormatInfo(3) Manual Page

## Name

VkPhysicalDeviceExternalImageFormatInfo - Structure specifying external
image creation parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the image capabilities compatible with an external memory
handle type, add a
[VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
structure to the `pNext` chain of the
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
structure and a `VkExternalImageFormatProperties` structure to the
`pNext` chain of the
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure.

The `VkPhysicalDeviceExternalImageFormatInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceExternalImageFormatInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalImageFormatInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_memory_capabilities
typedef VkPhysicalDeviceExternalImageFormatInfo VkPhysicalDeviceExternalImageFormatInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the memory handle type that will be used with the
  memory associated with the image.

## <a href="#_description" class="anchor"></a>Description

If `handleType` is 0,
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
will behave as if
[VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
was not present, and
[VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
will be ignored.

If `handleType` is not compatible with the `format`, `type`, `tiling`,
`usage`, and `flags` specified in
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html),
then
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
returns `VK_ERROR_FORMAT_NOT_SUPPORTED`.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceExternalImageFormatInfo-sType-sType"
  id="VUID-VkPhysicalDeviceExternalImageFormatInfo-sType-sType"></a>
  VUID-VkPhysicalDeviceExternalImageFormatInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO`

- <a
  href="#VUID-VkPhysicalDeviceExternalImageFormatInfo-handleType-parameter"
  id="VUID-VkPhysicalDeviceExternalImageFormatInfo-handleType-parameter"></a>
  VUID-VkPhysicalDeviceExternalImageFormatInfo-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExternalImageFormatInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
