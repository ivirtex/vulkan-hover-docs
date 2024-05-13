# VkInstanceCreateInfo(3) Manual Page

## Name

VkInstanceCreateInfo - Structure specifying parameters of a newly
created instance



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkInstanceCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkInstanceCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkInstanceCreateFlags       flags;
    const VkApplicationInfo*    pApplicationInfo;
    uint32_t                    enabledLayerCount;
    const char* const*          ppEnabledLayerNames;
    uint32_t                    enabledExtensionCount;
    const char* const*          ppEnabledExtensionNames;
} VkInstanceCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlagBits.html) indicating
  the behavior of the instance.

- `pApplicationInfo` is `NULL` or a pointer to a `VkApplicationInfo`
  structure. If not `NULL`, this information helps implementations
  recognize behavior inherent to classes of applications.
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html) is defined in detail
  below.

- `enabledLayerCount` is the number of global layers to enable.

- `ppEnabledLayerNames` is a pointer to an array of `enabledLayerCount`
  null-terminated UTF-8 strings containing the names of layers to enable
  for the created instance. The layers are loaded in the order they are
  listed in this array, with the first array element being the closest
  to the application, and the last array element being the closest to
  the driver. See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-layers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-layers</a>
  section for further details.

- `enabledExtensionCount` is the number of global extensions to enable.

- `ppEnabledExtensionNames` is a pointer to an array of
  `enabledExtensionCount` null-terminated UTF-8 strings containing the
  names of extensions to enable.

## <a href="#_description" class="anchor"></a>Description

To capture events that occur while creating or destroying an instance,
an application **can** link a
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)
structure or a
[VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)
structure to the `pNext` element of the `VkInstanceCreateInfo` structure
given to `vkCreateInstance`. This callback is only valid for the
duration of the [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html) and the
[vkDestroyInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyInstance.html) call. Use
[vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html) or
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html) to
create persistent callback objects.

An application can add additional drivers by including the
[VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html)
struct to the `pNext` element of the `VkInstanceCreateInfo` structure
given to `vkCreateInstance`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html">VkDirectDriverLoadingListLUNARG</a>
allows applications to ship drivers with themselves. Only drivers that
are designed to work with it should be used, such as drivers that
implement Vulkan in software or that implement Vulkan by translating it
to a different API. Any driver that requires installation should not be
used, such as hardware drivers.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkInstanceCreateInfo-pNext-04925"
  id="VUID-VkInstanceCreateInfo-pNext-04925"></a>
  VUID-VkInstanceCreateInfo-pNext-04925  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a
  `VkDebugReportCallbackCreateInfoEXT` structure, the list of enabled
  extensions in `ppEnabledExtensionNames` **must** contain
  [`VK_EXT_debug_report`](VK_EXT_debug_report.html)

- <a href="#VUID-VkInstanceCreateInfo-pNext-04926"
  id="VUID-VkInstanceCreateInfo-pNext-04926"></a>
  VUID-VkInstanceCreateInfo-pNext-04926  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a
  `VkDebugUtilsMessengerCreateInfoEXT` structure, the list of enabled
  extensions in `ppEnabledExtensionNames` **must** contain
  [`VK_EXT_debug_utils`](VK_EXT_debug_utils.html)

- <a href="#VUID-VkInstanceCreateInfo-pNext-06779"
  id="VUID-VkInstanceCreateInfo-pNext-06779"></a>
  VUID-VkInstanceCreateInfo-pNext-06779  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be either
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` or
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT`

- <a href="#VUID-VkInstanceCreateInfo-flags-06559"
  id="VUID-VkInstanceCreateInfo-flags-06559"></a>
  VUID-VkInstanceCreateInfo-flags-06559  
  If `flags` has the `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR`
  bit set, the list of enabled extensions in `ppEnabledExtensionNames`
  **must** contain
  [`VK_KHR_portability_enumeration`](VK_KHR_portability_enumeration.html)

- <a href="#VUID-VkInstanceCreateInfo-pNext-09400"
  id="VUID-VkInstanceCreateInfo-pNext-09400"></a>
  VUID-VkInstanceCreateInfo-pNext-09400  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a
  [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html)
  structure, the list of enabled extensions in `ppEnabledExtensionNames`
  **must** contain
  [VK_LUNARG_direct_driver_loading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUNARG_direct_driver_loading.html)

Valid Usage (Implicit)

- <a href="#VUID-VkInstanceCreateInfo-sType-sType"
  id="VUID-VkInstanceCreateInfo-sType-sType"></a>
  VUID-VkInstanceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO`

- <a href="#VUID-VkInstanceCreateInfo-pNext-pNext"
  id="VUID-VkInstanceCreateInfo-pNext-pNext"></a>
  VUID-VkInstanceCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html),
  [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html),
  [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html),
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html),
  [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html), or
  [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFlagsEXT.html)

- <a href="#VUID-VkInstanceCreateInfo-sType-unique"
  id="VUID-VkInstanceCreateInfo-sType-unique"></a>
  VUID-VkInstanceCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html),
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  or [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)

- <a href="#VUID-VkInstanceCreateInfo-flags-parameter"
  id="VUID-VkInstanceCreateInfo-flags-parameter"></a>
  VUID-VkInstanceCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlagBits.html) values

- <a href="#VUID-VkInstanceCreateInfo-pApplicationInfo-parameter"
  id="VUID-VkInstanceCreateInfo-pApplicationInfo-parameter"></a>
  VUID-VkInstanceCreateInfo-pApplicationInfo-parameter  
  If `pApplicationInfo` is not `NULL`, `pApplicationInfo` **must** be a
  valid pointer to a valid [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)
  structure

- <a href="#VUID-VkInstanceCreateInfo-ppEnabledLayerNames-parameter"
  id="VUID-VkInstanceCreateInfo-ppEnabledLayerNames-parameter"></a>
  VUID-VkInstanceCreateInfo-ppEnabledLayerNames-parameter  
  If `enabledLayerCount` is not `0`, `ppEnabledLayerNames` **must** be a
  valid pointer to an array of `enabledLayerCount` null-terminated UTF-8
  strings

- <a href="#VUID-VkInstanceCreateInfo-ppEnabledExtensionNames-parameter"
  id="VUID-VkInstanceCreateInfo-ppEnabledExtensionNames-parameter"></a>
  VUID-VkInstanceCreateInfo-ppEnabledExtensionNames-parameter  
  If `enabledExtensionCount` is not `0`, `ppEnabledExtensionNames`
  **must** be a valid pointer to an array of `enabledExtensionCount`
  null-terminated UTF-8 strings

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html),
[VkInstanceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInstanceCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
