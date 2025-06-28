# VkDeviceAddressBindingCallbackDataEXT(3) Manual Page

## Name

VkDeviceAddressBindingCallbackDataEXT - Structure specifying parameters returned to the callback



## [](#_c_specification)C Specification

The definition of `VkDeviceAddressBindingCallbackDataEXT` is:

```c++
// Provided by VK_EXT_device_address_binding_report
typedef struct VkDeviceAddressBindingCallbackDataEXT {
    VkStructureType                   sType;
    void*                             pNext;
    VkDeviceAddressBindingFlagsEXT    flags;
    VkDeviceAddress                   baseAddress;
    VkDeviceSize                      size;
    VkDeviceAddressBindingTypeEXT     bindingType;
} VkDeviceAddressBindingCallbackDataEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDeviceAddressBindingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingFlagBitsEXT.html) specifying additional information about the binding event that caused the callback to be called.
- `baseAddress` is a GPU-accessible virtual address identifying the start of a region of the virtual address space associated with a Vulkan object, as identified by the `pObjects` member of [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html).
- `size` is the size in bytes of a region of GPU-accessible virtual address space.
- `bindingType` is a [VkDeviceAddressBindingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingTypeEXT.html) specifying the type of binding event that caused the callback to be called.

## [](#_description)Description

If the [`reportAddressBinding`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-reportAddressBinding) feature is enabled and the implementation binds or unbinds a region of virtual address space associated with a Vulkan object, the implementation **must** submit a debug message with the following properties:

- `messageSeverity` equal to `VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT`
- `messageTypes` equal to `VK_DEBUG_UTILS_MESSAGE_TYPE_DEVICE_ADDRESS_BINDING_BIT_EXT`
- `VkDebugUtilsMessengerCallbackDataEXT`::`pObjects` **must** identify the associated Vulkan object
- `VkDeviceAddressBindingCallbackDataEXT` **must** be included in the `pNext` chain of `VkDebugUtilsMessengerCallbackDataEXT`

These debug messages **must** be emitted both for GPU virtual address space regions that are explicitly bound to a Vulkan object via the `vkBind`\*Memory/`vkBind`\*Memory2 functions, and for those that are implicitly generated via memory allocation or importing external memory.

An implementation **may** report binding events associated with a Vulkan object via `VkDebugUtilsMessengerEXT` prior to the object becoming visible to an application via other Vulkan commands. For example, object creation functions **may** report binding events that occur during an objects creation. In such cases, `VkDeviceAddressBindingCallbackDataEXT`::`flags` **must** include `VK_DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT`.

Object handles reported in this manner are not [valid object handles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-validusage-handles), and **must** not be used as an input parameter to any Vulkan command.

Any valid object handle returned by an object creation function **must** match the handle specified via any previously reported binding events associated with the object’s creation.

Valid Usage (Implicit)

- [](#VUID-VkDeviceAddressBindingCallbackDataEXT-sType-sType)VUID-VkDeviceAddressBindingCallbackDataEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_ADDRESS_BINDING_CALLBACK_DATA_EXT`
- [](#VUID-VkDeviceAddressBindingCallbackDataEXT-flags-parameter)VUID-VkDeviceAddressBindingCallbackDataEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkDeviceAddressBindingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingFlagBitsEXT.html) values
- [](#VUID-VkDeviceAddressBindingCallbackDataEXT-bindingType-parameter)VUID-VkDeviceAddressBindingCallbackDataEXT-bindingType-parameter  
  `bindingType` **must** be a valid [VkDeviceAddressBindingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingTypeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_address\_binding\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_address_binding_report.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceAddressBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingFlagsEXT.html), [VkDeviceAddressBindingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingTypeEXT.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceAddressBindingCallbackDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0