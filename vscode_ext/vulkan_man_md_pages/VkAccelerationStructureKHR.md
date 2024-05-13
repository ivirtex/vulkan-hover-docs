# VkAccelerationStructureKHR(3) Manual Page

## Name

VkAccelerationStructureKHR - Opaque handle to an acceleration structure
object



## <a href="#_c_specification" class="anchor"></a>C Specification

Acceleration structures are opaque data structures that are built by the
implementation to more efficiently perform spatial queries on the
provided geometric data. For this extension, an acceleration structure
is either a top-level acceleration structure containing a set of
bottom-level acceleration structures or a bottom-level acceleration
structure containing either a set of axis-aligned bounding boxes for
custom geometry or a set of triangles.

Each instance in the top-level acceleration structure contains a
reference to a bottom-level acceleration structure as well as an
instance transform plus information required to index into the shader
bindings. The top-level acceleration structure is what is bound to the
acceleration descriptor, for example to trace inside the shader in the
ray tracing pipeline.

Acceleration structures are represented by `VkAccelerationStructureKHR`
handles:

``` c
// Provided by VK_KHR_acceleration_structure
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkAccelerationStructureKHR)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html),
[VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html),
[VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html),
[VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html),
[VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html),
[VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html),
[VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html),
[vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html),
[vkCreateAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureKHR.html),
[vkDestroyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyAccelerationStructureKHR.html),
[vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
