# StandaloneSpirv(3) Manual Page

## Name

StandaloneSpirv - Standalone SPIR-V Validation



## <a href="#_description" class="anchor"></a>Description

The following rules **can** be validated with only the SPIR-V module
itself. They do not depend on knowledge of the implementation and its
capabilities or knowledge of runtime information, such as enabled
features.

Valid Usage

- <a href="#VUID-StandaloneSpirv-None-04633"
  id="VUID-StandaloneSpirv-None-04633"></a>
  VUID-StandaloneSpirv-None-04633  
  Every entry point **must** have no return value and accept no
  arguments

- <a href="#VUID-StandaloneSpirv-None-04634"
  id="VUID-StandaloneSpirv-None-04634"></a>
  VUID-StandaloneSpirv-None-04634  
  The static function-call graph for an entry point **must** not contain
  cycles; that is, static recursion is not allowed

- <a href="#VUID-StandaloneSpirv-None-04635"
  id="VUID-StandaloneSpirv-None-04635"></a>
  VUID-StandaloneSpirv-None-04635  
  The `Logical` or `PhysicalStorageBuffer64` addressing model **must**
  be selected

- <a href="#VUID-StandaloneSpirv-None-04636"
  id="VUID-StandaloneSpirv-None-04636"></a>
  VUID-StandaloneSpirv-None-04636  
  `Scope` for execution **must** be limited to `Workgroup` or `Subgroup`

- <a href="#VUID-StandaloneSpirv-None-04637"
  id="VUID-StandaloneSpirv-None-04637"></a>
  VUID-StandaloneSpirv-None-04637  
  If the `Scope` for execution is `Workgroup`, then it **must** only be
  used in the task, mesh, tessellation control, or compute `Execution`
  `Model`

- <a href="#VUID-StandaloneSpirv-None-04638"
  id="VUID-StandaloneSpirv-None-04638"></a>
  VUID-StandaloneSpirv-None-04638  
  `Scope` for memory **must** be limited to `Device`, `QueueFamily`,
  `Workgroup`, `ShaderCallKHR`, `Subgroup`, or `Invocation`

- <a href="#VUID-StandaloneSpirv-ExecutionModel-07320"
  id="VUID-StandaloneSpirv-ExecutionModel-07320"></a>
  VUID-StandaloneSpirv-ExecutionModel-07320  
  If the `Execution` `Model` is `TessellationControl`, and the
  `MemoryModel` is `GLSL450`, the `Scope` for memory **must** not be
  `Workgroup`

- <a href="#VUID-StandaloneSpirv-None-07321"
  id="VUID-StandaloneSpirv-None-07321"></a>
  VUID-StandaloneSpirv-None-07321  
  If the `Scope` for memory is `Workgroup`, then it **must** only be
  used in the task, mesh, tessellation control, or compute `Execution`
  `Model`

- <a href="#VUID-StandaloneSpirv-None-04640"
  id="VUID-StandaloneSpirv-None-04640"></a>
  VUID-StandaloneSpirv-None-04640  
  If the `Scope` for memory is `ShaderCallKHR`, then it **must** only be
  used in ray generation, intersection, closest hit, any-hit, miss, and
  callable `Execution` `Model`

- <a href="#VUID-StandaloneSpirv-None-04641"
  id="VUID-StandaloneSpirv-None-04641"></a>
  VUID-StandaloneSpirv-None-04641  
  If the `Scope` for memory is `Invocation`, then memory semantics
  **must** be `None`

- <a href="#VUID-StandaloneSpirv-None-04642"
  id="VUID-StandaloneSpirv-None-04642"></a>
  VUID-StandaloneSpirv-None-04642  
  `Scope` for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
  target="_blank" rel="noopener">group operations</a> **must** be
  limited to `Subgroup`

- <a href="#VUID-StandaloneSpirv-SubgroupVoteKHR-07951"
  id="VUID-StandaloneSpirv-SubgroupVoteKHR-07951"></a>
  VUID-StandaloneSpirv-SubgroupVoteKHR-07951  
  If none of the `SubgroupVoteKHR`, `GroupNonUniform`, or
  `SubgroupBallotKHR` capabilities are declared, `Scope` for memory
  **must** not be `Subgroup`

- <a href="#VUID-StandaloneSpirv-None-04643"
  id="VUID-StandaloneSpirv-None-04643"></a>
  VUID-StandaloneSpirv-None-04643  
  `Storage` `Class` **must** be limited to `UniformConstant`, `Input`,
  `Uniform`, `Output`, `Workgroup`, `Private`, `Function`,
  `PushConstant`, `Image`, `StorageBuffer`, `RayPayloadKHR`,
  `IncomingRayPayloadKHR`, `HitAttributeKHR`, `CallableDataKHR`,
  `IncomingCallableDataKHR`, `ShaderRecordBufferKHR`,
  `PhysicalStorageBuffer`, or `TileImageEXT`

- <a href="#VUID-StandaloneSpirv-None-04644"
  id="VUID-StandaloneSpirv-None-04644"></a>
  VUID-StandaloneSpirv-None-04644  
  If the `Storage` `Class` is `Output`, then it **must** not be used in
  the `GlCompute`, `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`,
  `ClosestHitKHR`, `MissKHR`, or `CallableKHR` `Execution` `Model`

- <a href="#VUID-StandaloneSpirv-None-04645"
  id="VUID-StandaloneSpirv-None-04645"></a>
  VUID-StandaloneSpirv-None-04645  
  If the `Storage` `Class` is `Workgroup`, then it **must** only be used
  in the task, mesh, or compute `Execution` `Model`

- <a href="#VUID-StandaloneSpirv-None-08720"
  id="VUID-StandaloneSpirv-None-08720"></a>
  VUID-StandaloneSpirv-None-08720  
  If the `Storage` `Class` is `TileImageEXT`, then it **must** only be
  used in the fragment execution model

- <a href="#VUID-StandaloneSpirv-OpAtomicStore-04730"
  id="VUID-StandaloneSpirv-OpAtomicStore-04730"></a>
  VUID-StandaloneSpirv-OpAtomicStore-04730  
  `OpAtomicStore` **must** not use `Acquire`, `AcquireRelease`, or
  `SequentiallyConsistent` memory semantics

- <a href="#VUID-StandaloneSpirv-OpAtomicLoad-04731"
  id="VUID-StandaloneSpirv-OpAtomicLoad-04731"></a>
  VUID-StandaloneSpirv-OpAtomicLoad-04731  
  `OpAtomicLoad` **must** not use `Release`, `AcquireRelease`, or
  `SequentiallyConsistent` memory semantics

- <a href="#VUID-StandaloneSpirv-OpMemoryBarrier-04732"
  id="VUID-StandaloneSpirv-OpMemoryBarrier-04732"></a>
  VUID-StandaloneSpirv-OpMemoryBarrier-04732  
  `OpMemoryBarrier` **must** use one of `Acquire`, `Release`,
  `AcquireRelease`, or `SequentiallyConsistent` memory semantics

- <a href="#VUID-StandaloneSpirv-OpMemoryBarrier-04733"
  id="VUID-StandaloneSpirv-OpMemoryBarrier-04733"></a>
  VUID-StandaloneSpirv-OpMemoryBarrier-04733  
  `OpMemoryBarrier` **must** include at least one `Storage` `Class`

- <a href="#VUID-StandaloneSpirv-OpControlBarrier-04650"
  id="VUID-StandaloneSpirv-OpControlBarrier-04650"></a>
  VUID-StandaloneSpirv-OpControlBarrier-04650  
  If the semantics for `OpControlBarrier` includes one of `Acquire`,
  `Release`, `AcquireRelease`, or `SequentiallyConsistent` memory
  semantics, then it **must** include at least one `Storage` `Class`

- <a href="#VUID-StandaloneSpirv-OpVariable-04651"
  id="VUID-StandaloneSpirv-OpVariable-04651"></a>
  VUID-StandaloneSpirv-OpVariable-04651  
  Any `OpVariable` with an `Initializer` operand **must** have `Output`,
  `Private`, `Function`, or `Workgroup` as its `Storage` `Class` operand

- <a href="#VUID-StandaloneSpirv-OpVariable-04734"
  id="VUID-StandaloneSpirv-OpVariable-04734"></a>
  VUID-StandaloneSpirv-OpVariable-04734  
  Any `OpVariable` with an `Initializer` operand and `Workgroup` as its
  `Storage` `Class` operand **must** use `OpConstantNull` as the
  initializer

- <a href="#VUID-StandaloneSpirv-OpReadClockKHR-04652"
  id="VUID-StandaloneSpirv-OpReadClockKHR-04652"></a>
  VUID-StandaloneSpirv-OpReadClockKHR-04652  
  `Scope` for `OpReadClockKHR` **must** be limited to `Subgroup` or
  `Device`

- <a href="#VUID-StandaloneSpirv-OriginLowerLeft-04653"
  id="VUID-StandaloneSpirv-OriginLowerLeft-04653"></a>
  VUID-StandaloneSpirv-OriginLowerLeft-04653  
  The `OriginLowerLeft` `Execution` `Mode` **must** not be used;
  fragment entry points **must** declare `OriginUpperLeft`

- <a href="#VUID-StandaloneSpirv-PixelCenterInteger-04654"
  id="VUID-StandaloneSpirv-PixelCenterInteger-04654"></a>
  VUID-StandaloneSpirv-PixelCenterInteger-04654  
  The `PixelCenterInteger` `Execution` `Mode` **must** not be used
  (pixels are always centered at half-integer coordinates)

- <a href="#VUID-StandaloneSpirv-UniformConstant-04655"
  id="VUID-StandaloneSpirv-UniformConstant-04655"></a>
  VUID-StandaloneSpirv-UniformConstant-04655  
  Any variable in the `UniformConstant` `Storage` `Class` **must** be
  typed as either `OpTypeImage`, `OpTypeSampler`, `OpTypeSampledImage`,
  `OpTypeAccelerationStructureKHR`, or an array of one of these types

- <a href="#VUID-StandaloneSpirv-Uniform-06807"
  id="VUID-StandaloneSpirv-Uniform-06807"></a>
  VUID-StandaloneSpirv-Uniform-06807  
  Any variable in the `Uniform` or `StorageBuffer` `Storage` `Class`
  **must** be typed as `OpTypeStruct` or an array of this type

- <a href="#VUID-StandaloneSpirv-PushConstant-06808"
  id="VUID-StandaloneSpirv-PushConstant-06808"></a>
  VUID-StandaloneSpirv-PushConstant-06808  
  Any variable in the `PushConstant` `Storage` `Class` **must** be typed
  as `OpTypeStruct`

- <a href="#VUID-StandaloneSpirv-OpTypeImage-04656"
  id="VUID-StandaloneSpirv-OpTypeImage-04656"></a>
  VUID-StandaloneSpirv-OpTypeImage-04656  
  `OpTypeImage` **must** declare a scalar 32-bit float, 64-bit integer,
  or 32-bit integer type for the “Sampled Type” (`RelaxedPrecision`
  **can** be applied to a sampling instruction and to the variable
  holding the result of a sampling instruction)

- <a href="#VUID-StandaloneSpirv-OpTypeImage-04657"
  id="VUID-StandaloneSpirv-OpTypeImage-04657"></a>
  VUID-StandaloneSpirv-OpTypeImage-04657  
  `OpTypeImage` **must** have a “Sampled” operand of 1 (sampled image)
  or 2 (storage image)

- <a href="#VUID-StandaloneSpirv-OpTypeSampledImage-06671"
  id="VUID-StandaloneSpirv-OpTypeSampledImage-06671"></a>
  VUID-StandaloneSpirv-OpTypeSampledImage-06671  
  `OpTypeSampledImage` **must** have a `OpTypeImage` with a “Sampled”
  operand of 1 (sampled image)

- <a href="#VUID-StandaloneSpirv-Image-04965"
  id="VUID-StandaloneSpirv-Image-04965"></a>
  VUID-StandaloneSpirv-Image-04965  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirv-type"
  target="_blank" rel="noopener">SPIR-V Type</a> of the `Image` `Format`
  operand of an `OpTypeImage` **must** match the `Sampled` `Type`, as
  defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-format-type-matching"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-format-type-matching</a>

- <a href="#VUID-StandaloneSpirv-OpImageTexelPointer-04658"
  id="VUID-StandaloneSpirv-OpImageTexelPointer-04658"></a>
  VUID-StandaloneSpirv-OpImageTexelPointer-04658  
  If an `OpImageTexelPointer` is used in an atomic operation, the image
  type of the `image` parameter to `OpImageTexelPointer` **must** have
  an image format of `R64i`, `R64ui`, `R32f`, `R32i`, or `R32ui`

- <a href="#VUID-StandaloneSpirv-OpImageQuerySizeLod-04659"
  id="VUID-StandaloneSpirv-OpImageQuerySizeLod-04659"></a>
  VUID-StandaloneSpirv-OpImageQuerySizeLod-04659  
  `OpImageQuerySizeLod`, `OpImageQueryLod`, and `OpImageQueryLevels`
  **must** only consume an “Image” operand whose type has its “Sampled”
  operand set to 1

- <a href="#VUID-StandaloneSpirv-OpTypeImage-09638"
  id="VUID-StandaloneSpirv-OpTypeImage-09638"></a>
  VUID-StandaloneSpirv-OpTypeImage-09638  
  An `OpTypeImage` **must** not have a “Dim” operand of `Rect`

- <a href="#VUID-StandaloneSpirv-OpTypeImage-06214"
  id="VUID-StandaloneSpirv-OpTypeImage-06214"></a>
  VUID-StandaloneSpirv-OpTypeImage-06214  
  An `OpTypeImage` with a “Dim” operand of `SubpassData` **must** have
  an “Arrayed” operand of 0 (non-arrayed) and a “Sampled” operand of 2
  (storage image)

- <a href="#VUID-StandaloneSpirv-SubpassData-04660"
  id="VUID-StandaloneSpirv-SubpassData-04660"></a>
  VUID-StandaloneSpirv-SubpassData-04660  
  The (u,v) coordinates used for a `SubpassData` **must** be the \<id\>
  of a constant vector (0,0).

- <a href="#VUID-StandaloneSpirv-OpTypeImage-06924"
  id="VUID-StandaloneSpirv-OpTypeImage-06924"></a>
  VUID-StandaloneSpirv-OpTypeImage-06924  
  Objects of types `OpTypeImage`, `OpTypeSampler`, `OpTypeSampledImage`,
  `OpTypeAccelerationStructureKHR`, and arrays of these types **must**
  not be stored to or modified

- <a href="#VUID-StandaloneSpirv-Uniform-06925"
  id="VUID-StandaloneSpirv-Uniform-06925"></a>
  VUID-StandaloneSpirv-Uniform-06925  
  Any variable in the `Uniform` `Storage` `Class` decorated as `Block`
  **must** not be stored to or modified

- <a href="#VUID-StandaloneSpirv-Offset-04663"
  id="VUID-StandaloneSpirv-Offset-04663"></a>
  VUID-StandaloneSpirv-Offset-04663  
  Image operand `Offset` **must** only be used with `OpImage*Gather`
  instructions

- <a href="#VUID-StandaloneSpirv-Offset-04865"
  id="VUID-StandaloneSpirv-Offset-04865"></a>
  VUID-StandaloneSpirv-Offset-04865  
  Any image instruction which uses an `Offset`, `ConstOffset`, or
  `ConstOffsets` image operand, must only consume a “Sampled Image”
  operand whose type has its “Sampled” operand set to 1

- <a href="#VUID-StandaloneSpirv-OpImageGather-04664"
  id="VUID-StandaloneSpirv-OpImageGather-04664"></a>
  VUID-StandaloneSpirv-OpImageGather-04664  
  The “Component” operand of `OpImageGather`, and `OpImageSparseGather`
  **must** be the \<id\> of a constant instruction

- <a href="#VUID-StandaloneSpirv-OpImage-04777"
  id="VUID-StandaloneSpirv-OpImage-04777"></a>
  VUID-StandaloneSpirv-OpImage-04777  
  `OpImage*Dref*` instructions **must** not consume an image whose `Dim`
  is 3D

- <a href="#VUID-StandaloneSpirv-None-04667"
  id="VUID-StandaloneSpirv-None-04667"></a>
  VUID-StandaloneSpirv-None-04667  
  Structure types **must** not contain opaque types

- <a href="#VUID-StandaloneSpirv-BuiltIn-04668"
  id="VUID-StandaloneSpirv-BuiltIn-04668"></a>
  VUID-StandaloneSpirv-BuiltIn-04668  
  Any `BuiltIn` decoration not listed in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables</a>
  **must** not be used

- <a href="#VUID-StandaloneSpirv-OpEntryPoint-09658"
  id="VUID-StandaloneSpirv-OpEntryPoint-09658"></a>
  VUID-StandaloneSpirv-OpEntryPoint-09658  
  For a given `OpEntryPoint`, any `BuiltIn` decoration **must** not be
  used more than once by the `Input` interface.

- <a href="#VUID-StandaloneSpirv-OpEntryPoint-09659"
  id="VUID-StandaloneSpirv-OpEntryPoint-09659"></a>
  VUID-StandaloneSpirv-OpEntryPoint-09659  
  For a given `OpEntryPoint`, any `BuiltIn` decoration **must** not be
  used more than once by the `Output` interface.

- <a href="#VUID-StandaloneSpirv-Location-06672"
  id="VUID-StandaloneSpirv-Location-06672"></a>
  VUID-StandaloneSpirv-Location-06672  
  The `Location` or `Component` decorations **must** only be used with
  the `Input`, `Output`, `RayPayloadKHR`, `IncomingRayPayloadKHR`,
  `HitAttributeKHR`, `HitObjectAttributeNV`, `CallableDataKHR`,
  `IncomingCallableDataKHR`, or `ShaderRecordBufferKHR` storage classes

- <a href="#VUID-StandaloneSpirv-Location-04915"
  id="VUID-StandaloneSpirv-Location-04915"></a>
  VUID-StandaloneSpirv-Location-04915  
  The `Location` or `Component` decorations **must** not be used with
  `BuiltIn`

- <a href="#VUID-StandaloneSpirv-Location-04916"
  id="VUID-StandaloneSpirv-Location-04916"></a>
  VUID-StandaloneSpirv-Location-04916  
  The `Location` decorations **must** be used on <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variables</a>

- <a href="#VUID-StandaloneSpirv-Location-04917"
  id="VUID-StandaloneSpirv-Location-04917"></a>
  VUID-StandaloneSpirv-Location-04917  
  If a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variable</a> is not a
  pointer to a `Block` decorated `OpTypeStruct`, then the `OpVariable`
  **must** have a `Location` decoration

- <a href="#VUID-StandaloneSpirv-Location-04918"
  id="VUID-StandaloneSpirv-Location-04918"></a>
  VUID-StandaloneSpirv-Location-04918  
  If a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variable</a> has a
  `Location` decoration, and the variable is a pointer to a
  `OpTypeStruct`, then the members of that structure **must** not have
  `Location` decorations

- <a href="#VUID-StandaloneSpirv-Location-04919"
  id="VUID-StandaloneSpirv-Location-04919"></a>
  VUID-StandaloneSpirv-Location-04919  
  If a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variable</a> does not have
  a `Location` decoration, and the variable is a pointer to a `Block`
  decorated `OpTypeStruct`, then each member of the struct **must** have
  a `Location` decoration

- <a href="#VUID-StandaloneSpirv-Component-04920"
  id="VUID-StandaloneSpirv-Component-04920"></a>
  VUID-StandaloneSpirv-Component-04920  
  The `Component` decoration value **must** not be greater than 3

- <a href="#VUID-StandaloneSpirv-Component-04921"
  id="VUID-StandaloneSpirv-Component-04921"></a>
  VUID-StandaloneSpirv-Component-04921  
  If the `Component` decoration is used on an `OpVariable` that has a
  `OpTypeVector` type with a `Component` `Type` with a `Width` that is
  less than or equal to 32, the sum of its `Component` `Count` and the
  `Component` decoration value **must** be less than or equal to 4

- <a href="#VUID-StandaloneSpirv-Component-04922"
  id="VUID-StandaloneSpirv-Component-04922"></a>
  VUID-StandaloneSpirv-Component-04922  
  If the `Component` decoration is used on an `OpVariable` that has a
  `OpTypeVector` type with a `Component` `Type` with a `Width` that is
  equal to 64, the sum of two times its `Component` `Count` and the
  `Component` decoration value **must** be less than or equal to 4

- <a href="#VUID-StandaloneSpirv-Component-04923"
  id="VUID-StandaloneSpirv-Component-04923"></a>
  VUID-StandaloneSpirv-Component-04923  
  The `Component` decorations value **must** not be 1 or 3 for scalar or
  two-component 64-bit data types

- <a href="#VUID-StandaloneSpirv-Component-04924"
  id="VUID-StandaloneSpirv-Component-04924"></a>
  VUID-StandaloneSpirv-Component-04924  
  The `Component` decorations **must** not be used with any type that is
  not a scalar or vector, or an array of such a type

- <a href="#VUID-StandaloneSpirv-Component-07703"
  id="VUID-StandaloneSpirv-Component-07703"></a>
  VUID-StandaloneSpirv-Component-07703  
  The `Component` decorations **must** not be used for a 64-bit vector
  type with more than two components

- <a href="#VUID-StandaloneSpirv-Input-09557"
  id="VUID-StandaloneSpirv-Input-09557"></a>
  VUID-StandaloneSpirv-Input-09557  
  The pointers of any `Input` or `Output` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">Interface user-defined variables</a>
  **must** not contain any `PhysicalStorageBuffer` `Storage` `Class`
  pointers

- <a href="#VUID-StandaloneSpirv-GLSLShared-04669"
  id="VUID-StandaloneSpirv-GLSLShared-04669"></a>
  VUID-StandaloneSpirv-GLSLShared-04669  
  The `GLSLShared` and `GLSLPacked` decorations **must** not be used

- <a href="#VUID-StandaloneSpirv-Flat-04670"
  id="VUID-StandaloneSpirv-Flat-04670"></a>
  VUID-StandaloneSpirv-Flat-04670  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations
  **must** only be used on variables with the `Output` or `Input`
  `Storage` `Class`

- <a href="#VUID-StandaloneSpirv-Flat-06201"
  id="VUID-StandaloneSpirv-Flat-06201"></a>
  VUID-StandaloneSpirv-Flat-06201  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations
  **must** not be used on variables with the `Output` storage class in a
  fragment shader

- <a href="#VUID-StandaloneSpirv-Flat-06202"
  id="VUID-StandaloneSpirv-Flat-06202"></a>
  VUID-StandaloneSpirv-Flat-06202  
  The `Flat`, `NoPerspective`, `Sample`, and `Centroid` decorations
  **must** not be used on variables with the `Input` storage class in a
  vertex shader

- <a href="#VUID-StandaloneSpirv-PerVertexKHR-06777"
  id="VUID-StandaloneSpirv-PerVertexKHR-06777"></a>
  VUID-StandaloneSpirv-PerVertexKHR-06777  
  The `PerVertexKHR` decoration **must** only be used on variables with
  the `Input` `Storage` `Class` in a fragment shader

- <a href="#VUID-StandaloneSpirv-Flat-04744"
  id="VUID-StandaloneSpirv-Flat-04744"></a>
  VUID-StandaloneSpirv-Flat-04744  
  Any variable with integer or double-precision floating-point type and
  with `Input` `Storage` `Class` in a fragment shader, **must** be
  decorated `Flat`

- <a href="#VUID-StandaloneSpirv-ViewportRelativeNV-04672"
  id="VUID-StandaloneSpirv-ViewportRelativeNV-04672"></a>
  VUID-StandaloneSpirv-ViewportRelativeNV-04672  
  The `ViewportRelativeNV` decoration **must** only be used on a
  variable decorated with `Layer` in the vertex, tessellation
  evaluation, or geometry shader stages

- <a href="#VUID-StandaloneSpirv-ViewportRelativeNV-04673"
  id="VUID-StandaloneSpirv-ViewportRelativeNV-04673"></a>
  VUID-StandaloneSpirv-ViewportRelativeNV-04673  
  The `ViewportRelativeNV` decoration **must** not be used unless a
  variable decorated with one of `ViewportIndex` or `ViewportMaskNV` is
  also statically used by the same `OpEntryPoint`

- <a href="#VUID-StandaloneSpirv-ViewportMaskNV-04674"
  id="VUID-StandaloneSpirv-ViewportMaskNV-04674"></a>
  VUID-StandaloneSpirv-ViewportMaskNV-04674  
  The `ViewportMaskNV` and `ViewportIndex` decorations **must** not both
  be statically used by one or more `OpEntryPoint`’s that form the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stages</a> of
  a graphics pipeline

- <a href="#VUID-StandaloneSpirv-FPRoundingMode-04675"
  id="VUID-StandaloneSpirv-FPRoundingMode-04675"></a>
  VUID-StandaloneSpirv-FPRoundingMode-04675  
  Rounding modes other than round-to-nearest-even and round-towards-zero
  **must** not be used for the `FPRoundingMode` decoration

- <a href="#VUID-StandaloneSpirv-Invariant-04677"
  id="VUID-StandaloneSpirv-Invariant-04677"></a>
  VUID-StandaloneSpirv-Invariant-04677  
  Variables decorated with `Invariant` and variables with structure
  types that have any members decorated with `Invariant` **must** be in
  the `Output` or `Input` `Storage` `Class`, `Invariant` used on an
  `Input` `Storage` `Class` variable or structure member has no effect

- <a href="#VUID-StandaloneSpirv-VulkanMemoryModel-04678"
  id="VUID-StandaloneSpirv-VulkanMemoryModel-04678"></a>
  VUID-StandaloneSpirv-VulkanMemoryModel-04678  
  <span id="builtin-volatile-semantics"></span> If the
  `VulkanMemoryModel` capability is not declared, the `Volatile`
  decoration **must** be used on any variable declaration that includes
  one of the `SMIDNV`, `WarpIDNV`, `SubgroupSize`,
  `SubgroupLocalInvocationId`, `SubgroupEqMask`, `SubgroupGeMask`,
  `SubgroupGtMask`, `SubgroupLeMask`, or `SubgroupLtMask` `BuiltIn`
  decorations when used in the ray generation, closest hit, miss,
  intersection, or callable shaders, or with the `RayTmaxKHR` `Builtin`
  decoration when used in an intersection shader

- <a href="#VUID-StandaloneSpirv-VulkanMemoryModel-04679"
  id="VUID-StandaloneSpirv-VulkanMemoryModel-04679"></a>
  VUID-StandaloneSpirv-VulkanMemoryModel-04679  
  If the `VulkanMemoryModel` capability is declared, the `OpLoad`
  instruction **must** use the `Volatile` memory semantics when it
  accesses into any variable that includes one of the `SMIDNV`,
  `WarpIDNV`, `SubgroupSize`, `SubgroupLocalInvocationId`,
  `SubgroupEqMask`, `SubgroupGeMask`, `SubgroupGtMask`,
  `SubgroupLeMask`, or `SubgroupLtMask` `BuiltIn` decorations when used
  in the ray generation, closest hit, miss, intersection, or callable
  shaders, or with the `RayTmaxKHR` `Builtin` decoration when used in an
  intersection shader

- <a href="#VUID-StandaloneSpirv-OpTypeRuntimeArray-04680"
  id="VUID-StandaloneSpirv-OpTypeRuntimeArray-04680"></a>
  VUID-StandaloneSpirv-OpTypeRuntimeArray-04680  
  `OpTypeRuntimeArray` **must** only be used for:

  - the last member of a `Block`-decorated `OpTypeStruct` in
    `StorageBuffer` or `PhysicalStorageBuffer` storage `Storage` `Class`

  - `BufferBlock`-decorated `OpTypeStruct` in the `Uniform` storage
    `Storage` `Class`

  - the outermost dimension of an arrayed variable in the
    `StorageBuffer`, `Uniform`, or `UniformConstant` storage `Storage`
    `Class`

  - variables in the `NodePayloadAMDX` storage `Storage` `Class` when
    the `CoalescingAMDX` `Execution` `Mode` is specified

- <a href="#VUID-StandaloneSpirv-Function-04681"
  id="VUID-StandaloneSpirv-Function-04681"></a>
  VUID-StandaloneSpirv-Function-04681  
  A type *T* that is an array sized with a specialization constant
  **must** neither be, nor be contained in, the type *T2* of a variable
  *V*, unless either: a) *T* is equal to *T2*, b) *V* is declared in the
  `Function`, or `Private` `Storage` `Class`, c) *V* is a non-Block
  variable in the `Workgroup` `Storage` `Class`, or d) *V* is an
  interface variable with an additional level of arrayness, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-matching"
  target="_blank" rel="noopener">as described in interface matching</a>,
  and *T* is the member type of the array type *T2*

- <a href="#VUID-StandaloneSpirv-OpControlBarrier-04682"
  id="VUID-StandaloneSpirv-OpControlBarrier-04682"></a>
  VUID-StandaloneSpirv-OpControlBarrier-04682  
  If `OpControlBarrier` is used in ray generation, intersection,
  any-hit, closest hit, miss, fragment, vertex, tessellation evaluation,
  or geometry shaders, the execution Scope **must** be `Subgroup`

- <a href="#VUID-StandaloneSpirv-LocalSize-06426"
  id="VUID-StandaloneSpirv-LocalSize-06426"></a>
  VUID-StandaloneSpirv-LocalSize-06426  
  For each compute shader entry point, either a `LocalSize` or
  `LocalSizeId` `Execution` `Mode`, or an object decorated with the
  `WorkgroupSize` decoration **must** be specified

- <a href="#VUID-StandaloneSpirv-DerivativeGroupQuadsNV-04684"
  id="VUID-StandaloneSpirv-DerivativeGroupQuadsNV-04684"></a>
  VUID-StandaloneSpirv-DerivativeGroupQuadsNV-04684  
  For compute shaders using the `DerivativeGroupQuadsNV` execution mode,
  the first two dimensions of the local workgroup size **must** be a
  multiple of two

- <a href="#VUID-StandaloneSpirv-DerivativeGroupLinearNV-04778"
  id="VUID-StandaloneSpirv-DerivativeGroupLinearNV-04778"></a>
  VUID-StandaloneSpirv-DerivativeGroupLinearNV-04778  
  For compute shaders using the `DerivativeGroupLinearNV` execution
  mode, the product of the dimensions of the local workgroup size
  **must** be a multiple of four

- <a href="#VUID-StandaloneSpirv-OpGroupNonUniformBallotBitCount-04685"
  id="VUID-StandaloneSpirv-OpGroupNonUniformBallotBitCount-04685"></a>
  VUID-StandaloneSpirv-OpGroupNonUniformBallotBitCount-04685  
  If `OpGroupNonUniformBallotBitCount` is used, the group operation
  **must** be limited to `Reduce`, `InclusiveScan`, or `ExclusiveScan`

- <a href="#VUID-StandaloneSpirv-None-04686"
  id="VUID-StandaloneSpirv-None-04686"></a>
  VUID-StandaloneSpirv-None-04686  
  The *Pointer* operand of all atomic instructions **must** have a
  `Storage` `Class` limited to `Uniform`, `Workgroup`, `Image`,
  `StorageBuffer`, `PhysicalStorageBuffer`, or `TaskPayloadWorkgroupEXT`

- <a href="#VUID-StandaloneSpirv-Offset-04687"
  id="VUID-StandaloneSpirv-Offset-04687"></a>
  VUID-StandaloneSpirv-Offset-04687  
  Output variables or block members decorated with `Offset` that have a
  64-bit type, or a composite type containing a 64-bit type, **must**
  specify an `Offset` value aligned to a 8 byte boundary

- <a href="#VUID-StandaloneSpirv-Offset-04689"
  id="VUID-StandaloneSpirv-Offset-04689"></a>
  VUID-StandaloneSpirv-Offset-04689  
  The size of any output block containing any member decorated with
  `Offset` that is a 64-bit type **must** be a multiple of 8

- <a href="#VUID-StandaloneSpirv-Offset-04690"
  id="VUID-StandaloneSpirv-Offset-04690"></a>
  VUID-StandaloneSpirv-Offset-04690  
  The first member of an output block specifying a `Offset` decoration
  **must** specify a `Offset` value that is aligned to an 8 byte
  boundary if that block contains any member decorated with `Offset` and
  is a 64-bit type

- <a href="#VUID-StandaloneSpirv-Offset-04691"
  id="VUID-StandaloneSpirv-Offset-04691"></a>
  VUID-StandaloneSpirv-Offset-04691  
  Output variables or block members decorated with `Offset` that have a
  32-bit type, or a composite type contains a 32-bit type, **must**
  specify an `Offset` value aligned to a 4 byte boundary

- <a href="#VUID-StandaloneSpirv-Offset-04692"
  id="VUID-StandaloneSpirv-Offset-04692"></a>
  VUID-StandaloneSpirv-Offset-04692  
  Output variables, blocks or block members decorated with `Offset`
  **must** only contain base types that have components that are either
  32-bit or 64-bit in size

- <a href="#VUID-StandaloneSpirv-Offset-04716"
  id="VUID-StandaloneSpirv-Offset-04716"></a>
  VUID-StandaloneSpirv-Offset-04716  
  Only variables or block members in the output interface decorated with
  `Offset` **can** be captured for transform feedback, and those
  variables or block members **must** also be decorated with `XfbBuffer`
  and `XfbStride`, or inherit `XfbBuffer` and `XfbStride` decorations
  from a block containing them

- <a href="#VUID-StandaloneSpirv-XfbBuffer-04693"
  id="VUID-StandaloneSpirv-XfbBuffer-04693"></a>
  VUID-StandaloneSpirv-XfbBuffer-04693  
  All variables or block members in the output interface of the entry
  point being compiled decorated with a specific `XfbBuffer` value
  **must** all be decorated with identical `XfbStride` values

- <a href="#VUID-StandaloneSpirv-Stream-04694"
  id="VUID-StandaloneSpirv-Stream-04694"></a>
  VUID-StandaloneSpirv-Stream-04694  
  If any variables or block members in the output interface of the entry
  point being compiled are decorated with `Stream`, then all variables
  belonging to the same `XfbBuffer` **must** specify the same `Stream`
  value

- <a href="#VUID-StandaloneSpirv-XfbBuffer-04696"
  id="VUID-StandaloneSpirv-XfbBuffer-04696"></a>
  VUID-StandaloneSpirv-XfbBuffer-04696  
  For any two variables or block members in the output interface of the
  entry point being compiled with the same `XfbBuffer` value, the ranges
  determined by the `Offset` decoration and the size of the type
  **must** not overlap

- <a href="#VUID-StandaloneSpirv-XfbBuffer-04697"
  id="VUID-StandaloneSpirv-XfbBuffer-04697"></a>
  VUID-StandaloneSpirv-XfbBuffer-04697  
  All block members in the output interface of the entry point being
  compiled that are in the same block and have a declared or inherited
  `XfbBuffer` decoration **must** specify the same `XfbBuffer` value

- <a href="#VUID-StandaloneSpirv-RayPayloadKHR-04698"
  id="VUID-StandaloneSpirv-RayPayloadKHR-04698"></a>
  VUID-StandaloneSpirv-RayPayloadKHR-04698  
  `RayPayloadKHR` `Storage` `Class` **must** only be used in ray
  generation, closest hit or miss shaders

- <a href="#VUID-StandaloneSpirv-IncomingRayPayloadKHR-04699"
  id="VUID-StandaloneSpirv-IncomingRayPayloadKHR-04699"></a>
  VUID-StandaloneSpirv-IncomingRayPayloadKHR-04699  
  `IncomingRayPayloadKHR` `Storage` `Class` **must** only be used in
  closest hit, any-hit, or miss shaders

- <a href="#VUID-StandaloneSpirv-IncomingRayPayloadKHR-04700"
  id="VUID-StandaloneSpirv-IncomingRayPayloadKHR-04700"></a>
  VUID-StandaloneSpirv-IncomingRayPayloadKHR-04700  
  There **must** be at most one variable with the
  `IncomingRayPayloadKHR` `Storage` `Class` in the input interface of an
  entry point

- <a href="#VUID-StandaloneSpirv-HitAttributeKHR-04701"
  id="VUID-StandaloneSpirv-HitAttributeKHR-04701"></a>
  VUID-StandaloneSpirv-HitAttributeKHR-04701  
  `HitAttributeKHR` `Storage` `Class` **must** only be used in
  intersection, any-hit, or closest hit shaders

- <a href="#VUID-StandaloneSpirv-HitAttributeKHR-04702"
  id="VUID-StandaloneSpirv-HitAttributeKHR-04702"></a>
  VUID-StandaloneSpirv-HitAttributeKHR-04702  
  There **must** be at most one variable with the `HitAttributeKHR`
  `Storage` `Class` in the input interface of an entry point

- <a href="#VUID-StandaloneSpirv-HitAttributeKHR-04703"
  id="VUID-StandaloneSpirv-HitAttributeKHR-04703"></a>
  VUID-StandaloneSpirv-HitAttributeKHR-04703  
  A variable with `HitAttributeKHR` `Storage` `Class` **must** only be
  written to in an intersection shader

- <a href="#VUID-StandaloneSpirv-CallableDataKHR-04704"
  id="VUID-StandaloneSpirv-CallableDataKHR-04704"></a>
  VUID-StandaloneSpirv-CallableDataKHR-04704  
  `CallableDataKHR` `Storage` `Class` **must** only be used in ray
  generation, closest hit, miss, and callable shaders

- <a href="#VUID-StandaloneSpirv-IncomingCallableDataKHR-04705"
  id="VUID-StandaloneSpirv-IncomingCallableDataKHR-04705"></a>
  VUID-StandaloneSpirv-IncomingCallableDataKHR-04705  
  `IncomingCallableDataKHR` `Storage` `Class` **must** only be used in
  callable shaders

- <a href="#VUID-StandaloneSpirv-IncomingCallableDataKHR-04706"
  id="VUID-StandaloneSpirv-IncomingCallableDataKHR-04706"></a>
  VUID-StandaloneSpirv-IncomingCallableDataKHR-04706  
  There **must** be at most one variable with the
  `IncomingCallableDataKHR` `Storage` `Class` in the input interface of
  an entry point

- <a href="#VUID-StandaloneSpirv-ShaderRecordBufferKHR-07119"
  id="VUID-StandaloneSpirv-ShaderRecordBufferKHR-07119"></a>
  VUID-StandaloneSpirv-ShaderRecordBufferKHR-07119  
  `ShaderRecordBufferKHR` `Storage` `Class` **must** only be used in ray
  generation, intersection, any-hit, closest hit, callable, or miss
  shaders

- <a href="#VUID-StandaloneSpirv-Base-07650"
  id="VUID-StandaloneSpirv-Base-07650"></a>
  VUID-StandaloneSpirv-Base-07650  
  The `Base` operand of `OpPtrAccessChain` **must** have a storage class
  of `Workgroup`, `StorageBuffer`, or `PhysicalStorageBuffer`

- <a href="#VUID-StandaloneSpirv-Base-07651"
  id="VUID-StandaloneSpirv-Base-07651"></a>
  VUID-StandaloneSpirv-Base-07651  
  If the `Base` operand of `OpPtrAccessChain` has a `Workgroup`
  `Storage` `Class`, then the `VariablePointers` capability **must** be
  declared

- <a href="#VUID-StandaloneSpirv-Base-07652"
  id="VUID-StandaloneSpirv-Base-07652"></a>
  VUID-StandaloneSpirv-Base-07652  
  If the `Base` operand of `OpPtrAccessChain` has a `StorageBuffer`
  `Storage` `Class`, then the `VariablePointers` or
  `VariablePointersStorageBuffer` capability **must** be declared

- <a href="#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04708"
  id="VUID-StandaloneSpirv-PhysicalStorageBuffer64-04708"></a>
  VUID-StandaloneSpirv-PhysicalStorageBuffer64-04708  
  If the `PhysicalStorageBuffer64` addressing model is enabled, all
  instructions that support memory access operands and that use a
  physical pointer **must** include the `Aligned` operand

- <a href="#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04709"
  id="VUID-StandaloneSpirv-PhysicalStorageBuffer64-04709"></a>
  VUID-StandaloneSpirv-PhysicalStorageBuffer64-04709  
  If the `PhysicalStorageBuffer64` addressing model is enabled, any
  access chain instruction that accesses into a `RowMajor` matrix
  **must** only be used as the `Pointer` operand to `OpLoad` or
  `OpStore`

- <a href="#VUID-StandaloneSpirv-PhysicalStorageBuffer64-04710"
  id="VUID-StandaloneSpirv-PhysicalStorageBuffer64-04710"></a>
  VUID-StandaloneSpirv-PhysicalStorageBuffer64-04710  
  If the `PhysicalStorageBuffer64` addressing model is enabled,
  `OpConvertUToPtr` and `OpConvertPtrToU` **must** use an integer type
  whose `Width` is 64

- <a href="#VUID-StandaloneSpirv-OpTypeForwardPointer-04711"
  id="VUID-StandaloneSpirv-OpTypeForwardPointer-04711"></a>
  VUID-StandaloneSpirv-OpTypeForwardPointer-04711  
  `OpTypeForwardPointer` **must** have a `Storage` `Class` of
  `PhysicalStorageBuffer`

- <a href="#VUID-StandaloneSpirv-None-04745"
  id="VUID-StandaloneSpirv-None-04745"></a>
  VUID-StandaloneSpirv-None-04745  
  All block members in a variable with a `Storage` `Class` of
  `PushConstant` declared as an array **must** only be accessed by
  dynamically uniform indices

- <a href="#VUID-StandaloneSpirv-OpVariable-06673"
  id="VUID-StandaloneSpirv-OpVariable-06673"></a>
  VUID-StandaloneSpirv-OpVariable-06673  
  There **must** not be more than one `OpVariable` in the `PushConstant`
  `Storage` `Class` listed in the `Interface` for each `OpEntryPoint`

- <a href="#VUID-StandaloneSpirv-OpEntryPoint-06674"
  id="VUID-StandaloneSpirv-OpEntryPoint-06674"></a>
  VUID-StandaloneSpirv-OpEntryPoint-06674  
  Each `OpEntryPoint` **must** not statically use more than one
  `OpVariable` in the `PushConstant` `Storage` `Class`

- <a href="#VUID-StandaloneSpirv-OpEntryPoint-08721"
  id="VUID-StandaloneSpirv-OpEntryPoint-08721"></a>
  VUID-StandaloneSpirv-OpEntryPoint-08721  
  Each `OpEntryPoint` **must** not have more than one `Input` variable
  assigned the same `Component` word inside a `Location` slot, either
  explicitly or implicitly

- <a href="#VUID-StandaloneSpirv-OpEntryPoint-08722"
  id="VUID-StandaloneSpirv-OpEntryPoint-08722"></a>
  VUID-StandaloneSpirv-OpEntryPoint-08722  
  Each `OpEntryPoint` **must** not have more than one `Output` variable
  assigned the same `Component` word inside a `Location` slot, either
  explicitly or implicitly

- <a href="#VUID-StandaloneSpirv-Result-04780"
  id="VUID-StandaloneSpirv-Result-04780"></a>
  VUID-StandaloneSpirv-Result-04780  
  The `Result` `Type` operand of any `OpImageRead` or
  `OpImageSparseRead` instruction **must** be a vector of four
  components

- <a href="#VUID-StandaloneSpirv-Base-04781"
  id="VUID-StandaloneSpirv-Base-04781"></a>
  VUID-StandaloneSpirv-Base-04781  
  The `Base` operand of any `OpBitCount`, `OpBitReverse`,
  `OpBitFieldInsert`, `OpBitFieldSExtract`, or `OpBitFieldUExtract`
  instruction **must** be a 32-bit integer scalar or a vector of 32-bit
  integers

- <a href="#VUID-StandaloneSpirv-PushConstant-06675"
  id="VUID-StandaloneSpirv-PushConstant-06675"></a>
  VUID-StandaloneSpirv-PushConstant-06675  
  Any variable in the `PushConstant` or `StorageBuffer` storage class
  **must** be decorated as `Block`

- <a href="#VUID-StandaloneSpirv-Uniform-06676"
  id="VUID-StandaloneSpirv-Uniform-06676"></a>
  VUID-StandaloneSpirv-Uniform-06676  
  Any variable in the `Uniform` `Storage` `Class` **must** be decorated
  as `Block` or `BufferBlock`

- <a href="#VUID-StandaloneSpirv-UniformConstant-06677"
  id="VUID-StandaloneSpirv-UniformConstant-06677"></a>
  VUID-StandaloneSpirv-UniformConstant-06677  
  Any variable in the `UniformConstant`, `StorageBuffer`, or `Uniform`
  `Storage` `Class` **must** be decorated with `DescriptorSet` and
  `Binding`

- <a href="#VUID-StandaloneSpirv-InputAttachmentIndex-06678"
  id="VUID-StandaloneSpirv-InputAttachmentIndex-06678"></a>
  VUID-StandaloneSpirv-InputAttachmentIndex-06678  
  Variables decorated with `InputAttachmentIndex` **must** be in the
  `UniformConstant` `Storage` `Class`

- <a href="#VUID-StandaloneSpirv-DescriptorSet-06491"
  id="VUID-StandaloneSpirv-DescriptorSet-06491"></a>
  VUID-StandaloneSpirv-DescriptorSet-06491  
  If a variable is decorated by `DescriptorSet` or `Binding`, the
  `Storage` `Class` **must** correspond to an entry in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-resources-storage-class-correspondence"
  target="_blank" rel="noopener">Shader Resource and Storage Class
  Correspondence</a>

- <a href="#VUID-StandaloneSpirv-Input-06778"
  id="VUID-StandaloneSpirv-Input-06778"></a>
  VUID-StandaloneSpirv-Input-06778  
  Variables with a `Storage` `Class` of `Input` in a fragment shader
  stage that are decorated with `PerVertexKHR` **must** be declared as
  arrays

- <a href="#VUID-StandaloneSpirv-MeshEXT-07102"
  id="VUID-StandaloneSpirv-MeshEXT-07102"></a>
  VUID-StandaloneSpirv-MeshEXT-07102  
  The module **must** not contain both an entry point that uses the
  `TaskEXT` or `MeshEXT` `Execution` `Model` and an entry point that
  uses the `TaskNV` or `MeshNV` `Execution` `Model`

- <a href="#VUID-StandaloneSpirv-MeshEXT-07106"
  id="VUID-StandaloneSpirv-MeshEXT-07106"></a>
  VUID-StandaloneSpirv-MeshEXT-07106  
  In mesh shaders using the `MeshEXT` `Execution` `Model`
  `OpSetMeshOutputsEXT` **must** be called before any outputs are
  written

- <a href="#VUID-StandaloneSpirv-MeshEXT-07107"
  id="VUID-StandaloneSpirv-MeshEXT-07107"></a>
  VUID-StandaloneSpirv-MeshEXT-07107  
  In mesh shaders using the `MeshEXT` `Execution` `Model` all variables
  declared as output **must** not be read from

- <a href="#VUID-StandaloneSpirv-MeshEXT-07108"
  id="VUID-StandaloneSpirv-MeshEXT-07108"></a>
  VUID-StandaloneSpirv-MeshEXT-07108  
  In mesh shaders using the `MeshEXT` `Execution` `Model` for
  `OpSetMeshOutputsEXT` instructions, the “Vertex Count” and “Primitive
  Count” operands **must** not depend on `ViewIndex`

- <a href="#VUID-StandaloneSpirv-MeshEXT-07109"
  id="VUID-StandaloneSpirv-MeshEXT-07109"></a>
  VUID-StandaloneSpirv-MeshEXT-07109  
  In mesh shaders using the `MeshEXT` `Execution` `Model` variables
  decorated with `PrimitivePointIndicesEXT`, `PrimitiveLineIndicesEXT`,
  or `PrimitiveTriangleIndicesEXT` declared as an array **must** not be
  accessed by indices that depend on `ViewIndex`

- <a href="#VUID-StandaloneSpirv-MeshEXT-07110"
  id="VUID-StandaloneSpirv-MeshEXT-07110"></a>
  VUID-StandaloneSpirv-MeshEXT-07110  
  In mesh shaders using the `MeshEXT` `Execution` `Model` any values
  stored in variables decorated with `PrimitivePointIndicesEXT`,
  `PrimitiveLineIndicesEXT`, or `PrimitiveTriangleIndicesEXT` **must**
  not depend on `ViewIndex`

- <a href="#VUID-StandaloneSpirv-MeshEXT-07111"
  id="VUID-StandaloneSpirv-MeshEXT-07111"></a>
  VUID-StandaloneSpirv-MeshEXT-07111  
  In mesh shaders using the `MeshEXT` `Execution` `Model` variables in
  workgroup or private `Storage` `Class` declared as or containing a
  composite type **must** not be accessed by indices that depend on
  `ViewIndex`

- <a href="#VUID-StandaloneSpirv-MeshEXT-07330"
  id="VUID-StandaloneSpirv-MeshEXT-07330"></a>
  VUID-StandaloneSpirv-MeshEXT-07330  
  In mesh shaders using the `MeshEXT` `Execution` `Model` the
  `OutputVertices` `Execution` `Mode` **must** be greater than 0

- <a href="#VUID-StandaloneSpirv-MeshEXT-07331"
  id="VUID-StandaloneSpirv-MeshEXT-07331"></a>
  VUID-StandaloneSpirv-MeshEXT-07331  
  In mesh shaders using the `MeshEXT` `Execution` `Model` the
  `OutputPrimitivesEXT` `Execution` `Mode` **must** be greater than 0

- <a href="#VUID-StandaloneSpirv-Input-07290"
  id="VUID-StandaloneSpirv-Input-07290"></a>
  VUID-StandaloneSpirv-Input-07290  
  Variables with a `Storage` `Class` of `Input` or `Output` and a type
  of `OpTypeBool` **must** be decorated with the `BuiltIn` decoration

- <a href="#VUID-StandaloneSpirv-TileImageEXT-08723"
  id="VUID-StandaloneSpirv-TileImageEXT-08723"></a>
  VUID-StandaloneSpirv-TileImageEXT-08723  
  The tile image variable declarations **must** obey the constraints on
  the `TileImageEXT` `Storage` `Class` and the `Location` decoration
  described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-fragmenttileimage"
  target="_blank" rel="noopener">Fragment Tile Image Interface</a>

- <a href="#VUID-StandaloneSpirv-None-08724"
  id="VUID-StandaloneSpirv-None-08724"></a>
  VUID-StandaloneSpirv-None-08724  
  The `TileImageEXT` `Storage` `Class` **must** only be used for
  declaring tile image variables

- <a href="#VUID-StandaloneSpirv-Pointer-08973"
  id="VUID-StandaloneSpirv-Pointer-08973"></a>
  VUID-StandaloneSpirv-Pointer-08973  
  The `Storage` `Class` of the `Pointer` operand to
  `OpCooperativeMatrixLoadKHR` or `OpCooperativeMatrixStoreKHR` **must**
  be limited to `Workgroup`, `StorageBuffer`, or `PhysicalStorageBuffer`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#StandaloneSpirv"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
