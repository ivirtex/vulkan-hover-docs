# VkPhysicalDeviceBufferDeviceAddressFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceBufferDeviceAddressFeaturesEXT - Structure describing
buffer address features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_buffer_device_address
typedef struct VkPhysicalDeviceBufferDeviceAddressFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           bufferDeviceAddress;
    VkBool32           bufferDeviceAddressCaptureReplay;
    VkBool32           bufferDeviceAddressMultiDevice;
} VkPhysicalDeviceBufferDeviceAddressFeaturesEXT;
```

``` c
// Provided by VK_EXT_buffer_device_address
typedef VkPhysicalDeviceBufferDeviceAddressFeaturesEXT VkPhysicalDeviceBufferAddressFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-bufferDeviceAddressEXT"></span>
  `bufferDeviceAddress` indicates that the implementation supports
  accessing buffer memory in shaders as storage buffers via an address
  queried from
  [vkGetBufferDeviceAddressEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddressEXT.html).

- <span id="features-bufferDeviceAddressCaptureReplayEXT"></span>
  `bufferDeviceAddressCaptureReplay` indicates that the implementation
  supports saving and reusing buffer addresses, e.g. for trace capture
  and replay.

- <span id="features-bufferDeviceAddressMultiDeviceEXT"></span>
  `bufferDeviceAddressMultiDevice` indicates that the implementation
  supports the `bufferDeviceAddress` feature for logical devices created
  with multiple physical devices. If this feature is not supported,
  buffer addresses **must** not be queried on a logical device created
  with more than one physical device.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>VkPhysicalDeviceBufferDeviceAddressFeaturesEXT</code>
structure has the same members as the
<code>VkPhysicalDeviceBufferDeviceAddressFeatures</code> structure, but
the functionality indicated by the members is expressed differently. The
features indicated by the
<code>VkPhysicalDeviceBufferDeviceAddressFeatures</code> structure
requires additional flags to be passed at memory allocation time, and
the capture and replay mechanism is built around opaque capture
addresses for buffer and memory objects.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceBufferDeviceAddressFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceBufferDeviceAddressFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceBufferDeviceAddressFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_buffer_device_address.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceBufferDeviceAddressFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
