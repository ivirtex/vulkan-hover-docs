# VkInstanceCreateInfo(3) Manual Page

## Name

VkInstanceCreateInfo - Structure specifying parameters of a newly created instance



## [](#_c_specification)C Specification

The `VkInstanceCreateInfo` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateFlagBits.html) indicating the behavior of the instance.
- `pApplicationInfo` is `NULL` or a pointer to a `VkApplicationInfo` structure. If not `NULL`, this information helps implementations recognize behavior inherent to classes of applications. [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html) is defined in detail below.
- `enabledLayerCount` is the number of global layers to enable.
- `ppEnabledLayerNames` is a pointer to an array of `enabledLayerCount` null-terminated UTF-8 strings containing the names of layers to enable for the created instance. The layers are loaded in the order they are listed in this array, with the first array element being the closest to the application, and the last array element being the closest to the driver. See the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-layers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-layers) section for further details.
- `enabledExtensionCount` is the number of global extensions to enable.
- `ppEnabledExtensionNames` is a pointer to an array of `enabledExtensionCount` null-terminated UTF-8 strings containing the names of extensions to enable.

## [](#_description)Description

To capture events that occur while creating or destroying an instance, an application **can** link a [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) structure or a [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html) structure to the `pNext` chain of the `VkInstanceCreateInfo` structure passed to `vkCreateInstance`. This callback is only valid for the duration of the [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html) and the [vkDestroyInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyInstance.html) call. Use [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html) or [vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugUtilsMessengerEXT.html) to create persistent callback objects.

An application can add additional drivers by including the [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html) structure in the `pNext` chain of the `VkInstanceCreateInfo` structure passed to `vkCreateInstance`.

Note

[VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html) allows applications to ship drivers with themselves. Only drivers that are designed to work with it should be used, such as drivers that implement Vulkan in software or that implement Vulkan by translating it to a different API. Any driver that requires installation should not be used, such as hardware drivers.

Valid Usage

- [](#VUID-VkInstanceCreateInfo-pNext-04925)VUID-VkInstanceCreateInfo-pNext-04925  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain `VK_EXT_debug_report`
- [](#VUID-VkInstanceCreateInfo-pNext-04926)VUID-VkInstanceCreateInfo-pNext-04926  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain `VK_EXT_debug_utils`
- [](#VUID-VkInstanceCreateInfo-pNext-06779)VUID-VkInstanceCreateInfo-pNext-06779  
  If the `pNext` chain includes a [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html) structure, its `exportObjectType` member **must** be either `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` or `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT`
- [](#VUID-VkInstanceCreateInfo-flags-06559)VUID-VkInstanceCreateInfo-flags-06559  
  If `flags` has the `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR` bit set, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain `VK_KHR_portability_enumeration`
- [](#VUID-VkInstanceCreateInfo-pNext-09400)VUID-VkInstanceCreateInfo-pNext-09400  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain [VK\_LUNARG\_direct\_driver\_loading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_LUNARG_direct_driver_loading.html)
- [](#VUID-VkInstanceCreateInfo-pNext-10242)VUID-VkInstanceCreateInfo-pNext-10242  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain [VK\_EXT\_layer\_settings](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_layer_settings.html)
- [](#VUID-VkInstanceCreateInfo-pNext-10243)VUID-VkInstanceCreateInfo-pNext-10243  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain [VK\_EXT\_validation\_features](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_features.html)
- [](#VUID-VkInstanceCreateInfo-pNext-10244)VUID-VkInstanceCreateInfo-pNext-10244  
  If the `pNext` chain of `VkInstanceCreateInfo` includes a [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html) structure, the list of enabled extensions in `ppEnabledExtensionNames` **must** contain [VK\_EXT\_validation\_flags](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_flags.html)

Valid Usage (Implicit)

- [](#VUID-VkInstanceCreateInfo-sType-sType)VUID-VkInstanceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO`
- [](#VUID-VkInstanceCreateInfo-pNext-pNext)VUID-VkInstanceCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html), [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html), [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html), [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html), [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html), [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html), or [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html)
- [](#VUID-VkInstanceCreateInfo-sType-unique)VUID-VkInstanceCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html), [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html), or [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html)
- [](#VUID-VkInstanceCreateInfo-flags-parameter)VUID-VkInstanceCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateFlagBits.html) values
- [](#VUID-VkInstanceCreateInfo-pApplicationInfo-parameter)VUID-VkInstanceCreateInfo-pApplicationInfo-parameter  
  If `pApplicationInfo` is not `NULL`, `pApplicationInfo` **must** be a valid pointer to a valid [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html) structure
- [](#VUID-VkInstanceCreateInfo-ppEnabledLayerNames-parameter)VUID-VkInstanceCreateInfo-ppEnabledLayerNames-parameter  
  If `enabledLayerCount` is not `0`, `ppEnabledLayerNames` **must** be a valid pointer to an array of `enabledLayerCount` null-terminated UTF-8 strings
- [](#VUID-VkInstanceCreateInfo-ppEnabledExtensionNames-parameter)VUID-VkInstanceCreateInfo-ppEnabledExtensionNames-parameter  
  If `enabledExtensionCount` is not `0`, `ppEnabledExtensionNames` **must** be a valid pointer to an array of `enabledExtensionCount` null-terminated UTF-8 strings

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html), [VkInstanceCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInstanceCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0