# VkPhysicalDeviceBufferDeviceAddressFeatures(3) Manual Page

## Name

VkPhysicalDeviceBufferDeviceAddressFeatures - Structure describing
buffer address features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceBufferDeviceAddressFeatures` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceBufferDeviceAddressFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           bufferDeviceAddress;
    VkBool32           bufferDeviceAddressCaptureReplay;
    VkBool32           bufferDeviceAddressMultiDevice;
} VkPhysicalDeviceBufferDeviceAddressFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_buffer_device_address
typedef VkPhysicalDeviceBufferDeviceAddressFeatures VkPhysicalDeviceBufferDeviceAddressFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-bufferDeviceAddress"></span>
  `bufferDeviceAddress` indicates that the implementation supports
  accessing buffer memory in shaders as storage buffers via an address
  queried from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html).

- <span id="extension-features-bufferDeviceAddressCaptureReplay"></span>
  `bufferDeviceAddressCaptureReplay` indicates that the implementation
  supports saving and reusing buffer and device addresses, e.g. for
  trace capture and replay.

- <span id="extension-features-bufferDeviceAddressMultiDevice"></span>
  `bufferDeviceAddressMultiDevice` indicates that the implementation
  supports the `bufferDeviceAddress` , `rayTracingPipeline` and
  `rayQuery` features for logical devices created with multiple physical
  devices. If this feature is not supported, buffer and acceleration
  structure addresses **must** not be queried on a logical device
  created with more than one physical device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>bufferDeviceAddressMultiDevice</code> exists to allow certain
legacy platforms to be able to support <code>bufferDeviceAddress</code>
without needing to support shared GPU virtual addresses for multi-device
configurations.</p></td>
</tr>
</tbody>
</table>

See [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html) for more
information.

If the `VkPhysicalDeviceBufferDeviceAddressFeatures` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceBufferDeviceAddressFeatures` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceBufferDeviceAddressFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceBufferDeviceAddressFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceBufferDeviceAddressFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceBufferDeviceAddressFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
