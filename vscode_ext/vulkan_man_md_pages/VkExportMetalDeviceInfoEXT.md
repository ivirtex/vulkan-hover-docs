# VkExportMetalDeviceInfoEXT(3) Manual Page

## Name

VkExportMetalDeviceInfoEXT - Structure that identifies a VkDevice object
and corresponding Metal MTLDevice object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export the Metal `MTLDevice` object underlying the
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) associated with a
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) object, include a `VkExportMetalDeviceInfoEXT`
structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalDeviceInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalDeviceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    MTLDevice_id       mtlDevice;
} VkExportMetalDeviceInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mtlDevice` is the Metal `id<MTLDevice>` object underlying the
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) associated with the
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) object identified in the call. The
  implementation will return the `MTLDevice` in this member, or it will
  return `NULL` if no `MTLDevice` could be found underlying the
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalDeviceInfoEXT-sType-sType"
  id="VUID-VkExportMetalDeviceInfoEXT-sType-sType"></a>
  VUID-VkExportMetalDeviceInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_DEVICE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalDeviceInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
